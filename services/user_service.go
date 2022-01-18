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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil { // before updating the user try to get the current one
		return nil, err
	}

	if err := current.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FistName != "" {
			current.FistName = user.FistName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FistName = user.FistName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	updateErr := current.Update()
	if updateErr != nil {
		return nil, updateErr
	}
	return current, nil
}
