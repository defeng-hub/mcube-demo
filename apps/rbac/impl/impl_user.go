package impl

import (
	"context"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/infraboard/mcube/exception"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateUser(ctx context.Context, req *rbac.CreateUserRequest) (*rbac.User, error) {
	user, err := rbac.NewUser(req)
	if err != nil {
		return nil, exception.NewBadRequest("validate create user error, %s", err)
	}

	//预加载sql
	stmt, err := s.db.Prepare(insertUserSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// 执行 sql
	//user_name, pwd, email, phone, address, state, user_type
	_, err = stmt.Exec(
		user.UserName, user.Pwd, user.Email,
		user.Phone, user.Address, user.State, user.UserType,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
func (s *service) QueryUser(context.Context, *rbac.QueryUserRequest) (*rbac.UserSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBook not implemented")
}
