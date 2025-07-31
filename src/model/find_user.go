package model

import (
	"github.com/diogoalvesf/my-routine-backend/src/configuration/logger"
	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) FindUser(userId string) (*UserDomain, *rest_err.RestErr) {
	logger.Info("FindUser method called", zap.String("journey", "findUser"))

	return nil, nil
}
