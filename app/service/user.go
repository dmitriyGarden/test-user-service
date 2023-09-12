package service

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dmitriyGarden/test-user-service/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type UserConfig interface {
	JWTSecret() []byte
}

type UserService struct {
	cfg     UserConfig
	storage model.IStorage
}

func New(cfg UserConfig, s model.IStorage) (*UserService, error) {
	return &UserService{
		cfg:     cfg,
		storage: s,
	}, nil
}

func (c *UserService) GetJWT(ctx context.Context, login, password string) (string, error) {
	user, err := c.getUser(ctx, login)
	if err != nil {
		return "", fmt.Errorf("getUser: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", model.ErrNotFound
	}
	res, err := c.generateJWTToken(user)
	if err != nil {
		return "", fmt.Errorf("generateJWTToken: %w", err)
	}
	return res, nil
}

func (c *UserService) GetUserFromToken(token string) (uuid.UUID, error) {
	claims := new(jwtClaim)
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("unexpected method %s. %w", token.Method.Alg(), model.ErrInvalidToken)
		}
		return c.cfg.JWTSecret(), nil
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("%v. %w", err, model.ErrInvalidToken)
	}
	if !t.Valid {
		return uuid.Nil, model.ErrInvalidToken
	}
	uid, err := uuid.Parse(claims.UID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%v. %w", err, model.ErrInvalidToken)
	}
	return uid, nil
}

type jwtClaim struct {
	UID string `json:"uid"`
	jwt.RegisteredClaims
}

func (c *UserService) generateJWTToken(user *model.UserData) (string, error) {
	now := jwt.NewNumericDate(time.Now())
	claims := &jwtClaim{
		UID: user.ID.String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			NotBefore: now,
			IssuedAt:  now,
			ID:        uuid.New().String(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := c.cfg.JWTSecret()
	fmt.Println(string(key))
	str, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("token.SignedString: %w", err)
	}
	return str, nil
}

func (c *UserService) getUser(ctx context.Context, email string) (*model.UserData, error) {
	return c.storage.GetUserByEmail(ctx, email)
}
