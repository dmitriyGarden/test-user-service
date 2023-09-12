package nats_server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dmitriyGarden/test-user-service/model"
	"github.com/google/uuid"
)

func (s *Server) GetBilling(ctx context.Context, uid uuid.UUID) (int64, error) {
	topic := subject{
		requestID: s.reqID(ctx),
		method:    transactionBalanceMethod,
	}
	payload, err := uid.MarshalBinary()
	if err != nil {
		return 0, fmt.Errorf("uid.MarshalBinary: %w", err)
	}
	str := topic.String()
	res, err := s.conn.RequestWithContext(ctx, str, payload)
	if err != nil {
		return 0, fmt.Errorf("conn.RequestWithContext: %w", err)
	}
	msg := new(respMessage)
	err = json.Unmarshal(res.Data, msg)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal: %w", err)
	}
	if msg.Type == errorMessage {
		text := ""
		err = json.Unmarshal(msg.Payload, &text)
		if err != nil {
			return 0, fmt.Errorf("unexpected error payload: %w", err)
		}
		return 0, fmt.Errorf("%s. %w", text, model.ErrInvalidResponse)
	}
	if msg.Type != successMessage {
		return 0, fmt.Errorf("unexpected message type:  %s. %w", msg.Type, model.ErrInvalidResponse)
	}
	amount := int64(0)
	err = json.Unmarshal(msg.Payload, &amount)
	if err != nil {
		return 0, fmt.Errorf("unexpected amount payload: %w", err)
	}
	return amount, nil
}
