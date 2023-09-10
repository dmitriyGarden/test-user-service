package service

import "context"

type UserConfig interface {
}

type UserService struct {
	cfg UserConfig
}

func New(cfg UserConfig) (*UserService, error) {
	return &UserService{cfg: cfg}, nil
}

func (c *UserService) GetJWT(_ context.Context, login, password string) (string, error) {
	return login + password, nil
}
