package model

import (
	"context"

	"github.com/google/uuid"
)

type IUser interface {
	GetJWT(ctx context.Context, login, password string) (string, error)
	GetUserFromToken(token string) (uuid.UUID, error)
}
