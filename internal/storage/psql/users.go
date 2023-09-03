package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	conn *sqlx.DB
}

func NewUserStorage(conn *sqlx.DB) *UserStorage {
	return &UserStorage{conn: conn}
}

type SelectUserParams struct {
	Id int
}

type User struct {
	ID int64 `db:"id"`
}

var ErrUserNotFound = errors.New("user not found")

func (s *UserStorage) GetUserByID(ctx context.Context, fields *SelectUserParams) (User, error) {
	var user User
	err := s.conn.GetContext(ctx, &user, "SELECT id FROM users WHERE id=?", fields.Id)
	if errors.Is(err, sql.ErrNoRows) {
		return user, fmt.Errorf("select error %w", ErrUserNotFound)
	} else if err != nil {
		return user, fmt.Errorf("select error %w", err)
	}

	return user, nil
}
