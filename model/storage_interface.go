package model

import "context"

type IStorage interface {
	GetUserByEmail(ctx context.Context, email string) (*UserData, error)
}
