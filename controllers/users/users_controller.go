package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// One way of implementation of JSON parsing
	// fmt.Println(user)
	// body, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	c.String(http.StatusBadRequest, "Unable to read request body")
	// 	return
	// }
	// if err := json.Unmarshal(body, &user); err != nil {
	// 	c.String(http.StatusBadRequest, "Invalid JSON Format")
	// 	return
	// }

	// Another way of implementation of JSON parsing
	if err := c.ShouldBindJSON(&user); err != nil {
		// return bad request to the caller
		restError := errors.NewBadRequestError("Invalid JSON Request Format")
		c.JSON(restError.Status, restError)
		return
	}

	result, saveError := services.CreateUser(user)

	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		c.String(http.StatusInternalServerError, "Unable to save user")
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not Implemented")
}
