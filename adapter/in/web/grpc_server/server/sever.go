package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/dmitriyGarden/test-user-service/adapter/in/web/grpc_server/api"
	"github.com/dmitriyGarden/test-user-service/model"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IConfig interface {
	GetListen() string
}

type GrpcServer struct {
	api.UnimplementedUserServer
	userService        model.IUser
	transactionService model.ITransaction
	cfg                IConfig
	l                  model.ILogger
	v                  *validator.Validate
}

func New(cnf IConfig, user model.IUser, tr model.ITransaction, l model.ILogger) (*GrpcServer, error) {
	return &GrpcServer{
		cfg:                cnf,
		userService:        user,
		transactionService: tr,
		l:                  l,
		v:                  validator.New(),
	}, nil
}

func (g *GrpcServer) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	req := struct {
		Login    string `validate:"required,email"`
		Password string `validate:"required"`
	}{
		request.GetLogin(),
		request.GetPassword(),
	}
	err := g.v.Struct(req)
	if err != nil {
		g.l.Warning("v.Struct: ", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	res, err := g.userService.GetJWT(ctx, req.Login, req.Password)
	if err != nil {
		g.l.Errorf("userService.GetJWT: %v", err)
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, "Invalid login or password")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &api.LoginResponse{Token: res}, nil
}

func (g *GrpcServer) Balance(ctx context.Context, _ *api.Empty) (*api.UserBalance, error) {
	uid, err := g.getUserId(ctx)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, model.ErrInvalidToken) || errors.Is(err, model.ErrAuthRequired) {
			code = codes.Unauthenticated
		}
		g.l.Errorf("getUserId: %v", err)
		return nil, status.Errorf(code, err.Error())
	}
	res, err := g.transactionService.GetBilling(g.setRequestID(ctx), uid)
	if err != nil {
		g.l.Errorf("GetBilling: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.UserBalance{Amount: res}, nil
}

func (g *GrpcServer) setRequestID(ctx context.Context) context.Context {
	reqID := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		headers := md.Get("x-request-id")
		if len(headers) > 0 {
			reqID = headers[0]
		}
	}
	if reqID == "" {
		reqID = uuid.NewString()
	}
	return context.WithValue(ctx, "requestID", reqID)
}

func (g *GrpcServer) getUserId(ctx context.Context) (uuid.UUID, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return uuid.Nil, model.ErrAuthRequired
	}
	headers := md.Get("authorization")
	if len(headers) == 0 {
		return uuid.Nil, model.ErrAuthRequired
	}
	token := strings.TrimPrefix(headers[0], "Bearer ")
	uid, err := g.userService.GetUserFromToken(token)
	if err != nil {
		return uuid.Nil, fmt.Errorf("userService.GetUserFromToken: %w", err)
	}
	return uid, nil
}

func (g *GrpcServer) Run(ctx context.Context) error {
	s := grpc.NewServer()
	api.RegisterUserServer(s, g)
	addr := g.cfg.GetListen()
	l, err := new(net.ListenConfig).Listen(ctx, "tcp", addr)
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}
	g.l.Infoln("Listen: ", addr)
	err = s.Serve(l)
	if err != nil {
		return fmt.Errorf("s.Serve: %w", err)
	}
	return nil
}
