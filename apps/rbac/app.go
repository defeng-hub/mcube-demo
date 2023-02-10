package rbac

import (
	"github.com/go-playground/validator/v10"
)

const (
	AppName = "rbac"
)

var (
	validate = validator.New()
)

func (req *CreateUserRequest) Validate() error {
	return validate.Struct(req)
}

func NewUser(req *CreateUserRequest) (*User, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	return &User{
		UserName: req.UserName,
		Pwd:      req.Pwd,
		Email:    req.Email,
		Phone:    req.Phone,
		Address:  req.Address,
		State:    1,
		UserType: "普通用户",
	}, nil
}
