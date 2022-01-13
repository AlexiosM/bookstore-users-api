package users

import (
	"fmt"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"
)

// data access object
// here is the entire logic to persist and retrieve
// this user from a given database

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get(userId int64) *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FistName = result.FistName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	user.Email = result.Email

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	usersDB[user.Id] = user

	return nil
}
