package model

import (
	"fmt"

	"github.com/HumCoding/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/HumCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
