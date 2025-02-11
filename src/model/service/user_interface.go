package service

import (
	"github.com/HunCoding/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/HunCoding/meu-primeiro-crud-go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

func (ud *userDomainService) deleteUser(string) *rest_err.RestErr {
	panic("unimplemented")
}

func (ud *userDomainService) FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	panic("unimplemented")
}

func (ud *userDomainService) UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr {
	panic("unimplemented")
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
}
