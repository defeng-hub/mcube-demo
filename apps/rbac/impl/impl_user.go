package impl

import (
	"context"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) CreateBook(context.Context, *rbac.CreateUserRequest) (*rbac.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBook not implemented")
}
func (s *service) QueryBook(context.Context, *rbac.QueryUserRequest) (*rbac.UserSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryBook not implemented")
}
