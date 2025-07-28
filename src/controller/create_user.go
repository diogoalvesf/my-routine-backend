package controller

import (
	"fmt"

	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
	"github.com/diogoalvesf/my-routine-backend/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("Invalid request body: %s", err.Error()))

		c.JSON(restErr.Code, restErr)

		return
	}
}
