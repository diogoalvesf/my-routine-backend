package model

import (
	"github.com/diogoalvesf/my-routine-backend/src/configuration/logger"
	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) DeleteUser(userId string) *rest_err.RestErr {
	logger.Info("DeleteUser method called", zap.String("journey", "deleteUser"))

	return nil
}
