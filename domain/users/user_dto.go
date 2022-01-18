package users

import (
	"strings"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"
)

// ==== data transfer object ====
// the object that we are going to be transferring
// from the persistance layer to the application and backword
type User struct {
	Id          int64  `json:"id"`
	FistName    string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string
}

func (user *User) Validate() *errors.RestErr {

	user.FistName = strings.TrimSpace(user.FistName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return &errors.RestErr{
			Message: "invalid email address",
		}
	}
	return nil
}
