package model

import "context"

type IUser interface {
	GetJWT(ctx context.Context, login, password string) (string, error)
}
