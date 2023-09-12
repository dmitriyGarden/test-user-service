package model

import (
	"context"

	"github.com/google/uuid"
)

type ITransaction interface {
	GetBilling(ctx context.Context, uid uuid.UUID) (int64, error)
}
