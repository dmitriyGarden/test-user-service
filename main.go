package main

import (
	"context"
	"log"

	"github.com/dmitriyGarden/test-user-service/adapter/in/web"
	"github.com/dmitriyGarden/test-user-service/adapter/out/persistence/pgres"
	"github.com/dmitriyGarden/test-user-service/adapter/out/transaction_service"
	"github.com/dmitriyGarden/test-user-service/app/service"
	"github.com/dmitriyGarden/test-user-service/pkg/config"
	"github.com/dmitriyGarden/test-user-service/pkg/logger"
	"github.com/nats-io/nats.go"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New: %v", err)
	}
	l := logger.New()
	storage, err := pgres.New(cfg)
	if err != nil {
		l.Fatal("pgres.New: ", err)
	}
	userService, err := service.New(cfg, storage)
	if err != nil {
		l.Fatal("service.New: ", err)
	}
	conn, err := nats.Connect(cfg.NatsConnectionString())
	if err != nil {
		l.Fatal("nats.Connect: ", err)
	}
	defer func(conn *nats.Conn) {
		_ = conn.Drain()
		conn.Close()
	}(conn)
	transaction, err := transaction_service.GetTransactionNatsAdapter(conn)
	if err != nil {
		l.Fatal("GetTransactionNatsAdapter: ", err)
	}
	webAdapter, err := web.GetWebGrpcAdapter(cfg, userService, transaction, l)
	if err != nil {
		l.Fatal("web.GetWebGrpcAdapter: ", err)
	}

	err = webAdapter.Run(context.Background())
	if err != nil {
		l.Fatal("webAdapter.Run: %v", err)
	}
	l.Infoln("Finished")
}
