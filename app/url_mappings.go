package app

import (
	ping "udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/ping_controller"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user", users.GetUser)
	router.POST("/users", users.CreateUser)
}
