package commands

import (
	logger "github.com/sirupsen/logrus"
)

type CreateUser struct{}

func NewCreateUser() CreateUser {
	return CreateUser{}
}

func (itself CreateUser) CreateUserCommand() {
	logger.Info("Command - Create Users")
	logger.Info("Command - Create Users")
	logger.Info("Command - Create Users")
}
