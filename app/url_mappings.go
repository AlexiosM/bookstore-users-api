package app

import (
	ping "udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/ping_controller"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)   // to modify a given record
	router.PATCH("/users/:user_id", users.UpdateUser) // used to partially update the table
}
