package users

import (
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/datasources/mysql/users_db"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/date_utils"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/mysql_utils"
)

// data access object
// here is the entire logic to persist and retrieve
// this user from a given database

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInstertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get(userId int64) *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow(user.Id)

	if getErr := row.Scan(&user.Id, &user.FistName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInstertUser)
	if err != nil {
		errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, saveErr := stmt.Exec(user.FistName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	// as an alternative to the above
	// insertResult, err := stmt.Exec(queryInstertUser, user.FistName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId

	return nil
}
