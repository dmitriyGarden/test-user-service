package pgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	sqr "github.com/Masterminds/squirrel"
	"github.com/dmitriyGarden/test-user-service/model"
)

const userTable = "userschema.user"

func (d *DB) GetUserByEmail(ctx context.Context, email string) (*model.UserData, error) {
	q, args, err := sqr.Select("id", "email", "password").
		From(userTable).
		Where(sqr.Eq{"email": email}).
		PlaceholderFormat(sqr.Dollar).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("ToSql: %w", err)
	}
	res := new(model.UserData)
	err = d.db.GetContext(ctx, res, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrNotFound
		}
		return nil, fmt.Errorf("db.Get: %w", err)
	}
	return res, nil
}
