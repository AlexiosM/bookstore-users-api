package services

import (
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/domain/users"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	return &user, nil
}
