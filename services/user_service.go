package services

import (
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/domain/users"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}

	if err := user.Get(userId); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	current := users.User{Id: user.Id}
	if err := current.Get(user.Id); err != nil { // before updating the user try to get the current one
		return nil, err
	}
	current.FistName = user.FistName
	current.LastName = user.LastName
	current.Email = user.Email

	updateErr := current.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return &current, nil
}
