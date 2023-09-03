package service

import (
	"context"
	"errors"
	storage "github.com/mc_transaction/internal/storage/psql"
)

type UserStorage interface {
	GetUserByID(ctx context.Context, fields *storage.SelectUserParams) (storage.User, error)
}

type UserService struct {
	storage UserStorage
}

func NewUser(storage UserStorage) *UserService {
	return &UserService{storage: storage}
}

type User struct {
	Id int64
}

var ErrUserNotFound = errors.New("user not found")

func (u *UserService) GetUserByID(ctx context.Context, id int) (User, error) {
	var user User
	userFromStor, err := u.storage.GetUserByID(ctx, &storage.SelectUserParams{
		Id: id,
	})
	if errors.Is(err, storage.ErrUserNotFound) {
		return user, ErrUserNotFound
	} else if err != nil {
		return user, err
	}
	user.Id = userFromStor.ID
	return user, nil
}
