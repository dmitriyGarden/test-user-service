package nats_server

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type Server struct {
	conn *nats.Conn
}

func New(conn *nats.Conn) (*Server, error) {
	return &Server{
		conn: conn,
	}, nil
}

func (s *Server) reqID(ctx context.Context) string {
	id, ok := ctx.Value("requestID").(string)
	if ok && id != "" {
		return id
	}
	return uuid.NewString()
}
