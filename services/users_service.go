package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
