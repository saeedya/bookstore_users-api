package users

import (
	"bookstore_users-api/utils/date"
	"bookstore_users-api/utils/errors"
	"fmt"
)

var userDB = make(map[int64]*User)

func (user *User) Get() *errors.RestError {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestError {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("cannot update user email, %s already in use", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("cannot update user %d, ID already exists", user.Id))
	}
	user.CreatedAt = date.GetNowString()

	userDB[user.Id] = user
	return nil
}
