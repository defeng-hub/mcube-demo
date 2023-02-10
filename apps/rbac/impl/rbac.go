package impl

import (
	"database/sql"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/defeng-hub/mcube-demo/conf"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	log logger.Logger
	rbac.UnimplementedUserServiceServer
}

func (s *service) Name() string {
	return rbac.AppName
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}
	s.db = db

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Registry(server *grpc.Server) {
	rbac.RegisterUserServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
