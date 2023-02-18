package impl

import (
	"context"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/defeng-hub/mcube-demo/util/exception"
	"github.com/infraboard/mcube/sqlbuilder"
)

func (s *service) CreateRole(ctx context.Context, req *rbac.CreateRoleRequest) (*rbac.Role, error) {
	role, err := rbac.NewRole(req)
	if err != nil {
		return nil, err
	}

	prepare, err := s.db.Prepare(insertRole)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()

	exec, err := prepare.Exec(role.GetRoleName(), role.GetDescription())
	if err != nil {
		return nil, err
	}

	role.RoleId, err = exec.LastInsertId()
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *service) DeleteRole(ctx context.Context, req *rbac.DeleteRoleRequest) (*rbac.Role, error) {
	if req.RoleId == 0 {
		newErr := exception.DefaultException(-1, "请求数据解析失败", nil)
		return nil, newErr
	}
	prepare, err := s.db.Prepare(deleteRole)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()

	exec, err := prepare.Exec(req.RoleId)
	if err != nil {
		return nil, err
	}

	affected, err := exec.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		newErr := exception.DefaultException(-5, "库中无此角色", nil)

		return nil, newErr
	}
	return nil, nil
}

func (s *service) QueryRole(ctx context.Context, req *rbac.QueryRoleRequest) (*rbac.RoleSet, error) {

	roleset := rbac.RoleSet{}
	sb := sqlbuilder.NewBuilder(selectRole)
	sb.Where("s_role.role_name LIKE ?", "%"+req.Keywords+"%")

	countSql, countArgs := sb.BuildCount()
	prepare, err := s.db.Prepare(countSql)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	err = prepare.QueryRow(countArgs...).Scan(&(roleset.Total))
	if err != nil {
		return nil, err
	}

	sql, args := sb.BuildQuery()
	stmt, err := s.db.Prepare(sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	query, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		role := rbac.Role{}
		err := query.Scan(&role.RoleId, &role.RoleName, &role.Description)
		if err != nil {
			return nil, err
		}
		roleset.Items = append(roleset.Items, &role)
	}
	return &roleset, nil
}

//func (s *service) Q
