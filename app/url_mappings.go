package app

import (
	ping "udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/ping_controller"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
}
