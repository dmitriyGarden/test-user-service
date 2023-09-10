package main

import (
	"context"
	"log"

	"github.com/dmitriyGarden/test-user-service/adapter/in/web"
	"github.com/dmitriyGarden/test-user-service/app/service"
	"github.com/dmitriyGarden/test-user-service/pkg/config"
	"github.com/dmitriyGarden/test-user-service/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New:")
	}
	l := logger.New()
	l.Debug(cfg)
	userService, err := service.New(cfg)
	if err != nil {
		l.Fatal("service.New: ", err)
	}
	webAdapter, err := web.GetWebGrpcAdapter(cfg, userService, l)
	if err != nil {
		l.Fatal("web.GetWebGrpcAdapter: ", err)
	}

	err = webAdapter.Run(context.Background())
	if err != nil {
		l.Fatal("webAdapter.Run: %v", err)
	}
	l.Info("Finished")
}
