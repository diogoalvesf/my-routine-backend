package controller

import (
	"net/http"

	"github.com/diogoalvesf/my-routine-backend/src/configuration/logger"
	"github.com/diogoalvesf/my-routine-backend/src/configuration/validation"
	"github.com/diogoalvesf/my-routine-backend/src/controller/model/request"
	"github.com/diogoalvesf/my-routine-backend/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)

		return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
	)

	if err := domain.CreateUser(); err != nil {
		if err := domain.CreateUser(); err != nil {
			c.JSON(err.Code, err)
			return
		}
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"),
	)

	c.String(http.StatusOK, "")
}
