package impl

import (
	"database/sql"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/defeng-hub/mcube-demo/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	db  *sql.DB
	gdb *gorm.DB
	log logger.Logger
	rbac.UnimplementedUserServiceServer
	rbac.UnimplementedRoleServiceServer
	rbac.UnimplementedUserRoleServiceServer
}

func (s *service) Name() string {
	return rbac.AppName
}

func (s *service) Config() (err error) {
	s.db, err = conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.gdb, err = conf.C().MySQL.GetGDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named(s.Name())
	return nil
}

func (s *service) Registry(server *grpc.Server) {
	rbac.RegisterUserServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
