package model

import "github.com/google/uuid"

type UserData struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
}
