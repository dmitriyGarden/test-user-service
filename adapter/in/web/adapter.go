package web

import (
	"github.com/dmitriyGarden/test-user-service/adapter/in/web/grpc_server/server"
	"github.com/dmitriyGarden/test-user-service/model"
)

func GetWebGrpcAdapter(cfg server.IConfig, srv model.IUser, tr model.ITransaction, l model.ILogger) (model.IWebAdapter, error) {
	return server.New(cfg, srv, tr, l)
}
