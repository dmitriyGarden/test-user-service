package transaction_service

import (
	"github.com/dmitriyGarden/test-user-service/adapter/out/transaction_service/nats_server"
	"github.com/dmitriyGarden/test-user-service/model"
	"github.com/nats-io/nats.go"
)

func GetTransactionNatsAdapter(conn *nats.Conn) (model.ITransaction, error) {
	return nats_server.New(conn)
}
