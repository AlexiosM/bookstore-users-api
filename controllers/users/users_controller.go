package users

import (
	"net/http"
	"strconv"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/domain/users"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/services"
	"udemy_code/StepByStepRESTMicroservices/bookstore-users-api/utils/errors"

	"github.com/gin-gonic/gin"
)

// provide functionality for the endpoints
// to interact against the users API
// all requests will be handled by this package.
// Take the request check all the parameters that we need to handle
// and send this handling to the service package

func CreateUser(c *gin.Context) {
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// fmt.Println(string(bytes))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// //	fmt.Println(string(bytes))
	// fmt.Println(user)
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user Id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}
