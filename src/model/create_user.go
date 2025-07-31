package model

import (
	"github.com/diogoalvesf/my-routine-backend/src/configuration/logger"
	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("CreateUser method called", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	return nil
}
