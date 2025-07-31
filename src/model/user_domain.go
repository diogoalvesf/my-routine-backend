package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/diogoalvesf/my-routine-backend/src/configuration/rest_err"
)

func NewUserDomain(name, email, password string) UserDomainInterface {
	return &UserDomain{
		name, email, password,
	}
}

type UserDomain struct {
	Name     string
	Email    string
	Password string
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()

	defer hash.Reset()

	hash.Write([]byte(ud.Password))

	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (*UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
