package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// provide functionality for the endpoints
// to interact against the users API
// all requests will be handled by this package

// take the request check all the parameters that we need to handle
// and send this handling to the service package
func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")

}

func CreateUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")

}
