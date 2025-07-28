package controller

import (
	"github.com/diogoalvesf/my-routine-backend/src/configuration/validation"
	"github.com/diogoalvesf/my-routine-backend/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}
}
