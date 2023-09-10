package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dmitriyGarden/test-user-service/adapter/in/web/grpc_server/api"
	"github.com/dmitriyGarden/test-user-service/model"
)

type IConfig interface {
	GetListen() string
}

type GrpcServer struct {
	api.UnimplementedUserServer
	userService model.IUser
	cfg         IConfig
	l           model.ILogger
}

func New(cnf IConfig, user model.IUser, l model.ILogger) (*GrpcServer, error) {
	return &GrpcServer{
		cfg:         cnf,
		userService: user,
		l:           l,
	}, nil
}

func (g *GrpcServer) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	res, err := g.userService.GetJWT(ctx, request.GetLogin(), request.GetPassword())
	if err != nil {
		g.l.Errorf("userService.GetJWT: %v", err)
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.InvalidArgument, "Invalid login or password")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	g.l.Debug(res)
	return &api.LoginResponse{Token: res}, nil
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
