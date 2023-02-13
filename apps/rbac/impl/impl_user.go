package impl

import (
	"context"
	"fmt"
	"github.com/defeng-hub/mcube-demo/apps/rbac"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/sqlbuilder"
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
	res, err := stmt.Exec(
		user.UserName, user.Pwd, user.Email,
		user.Phone, user.Address, user.State, user.UserType,
	)
	if err != nil {
		return nil, err
	}
	user.UserId, _ = res.LastInsertId()
	return user, nil
}

func (s *service) QueryUser(ctx context.Context, req *rbac.QueryUserRequest) (*rbac.UserSet, error) {
	userSet, _ := rbac.NewUserSet()

	query := sqlbuilder.NewBuilder(QueryUsersSQL)
	//添加关键字查询
	if req.Keywords != "" {
		query.Where("s_user.user_name LIKE ? or s_user.email LIKE ? or s_user.phone LIKE ? or s_user.address LIKE ?",
			"%"+req.Keywords+"%", "%"+req.Keywords+"%",
			"%"+req.Keywords+"%", "%"+req.Keywords+"%")
	}

	// 添加分页
	query.Limit(int64((req.Page.PageNumber-1)*req.Page.PageSize), uint(req.Page.PageSize))

	// 计算 total总数
	countSql, countArgs := query.BuildCount()
	countStmt, err2 := s.db.Prepare(countSql)
	if err2 != nil {
		return nil, err2
	}
	defer countStmt.Close()
	countStmt.QueryRow(countArgs...).Scan(&userSet.Total)

	// 建立查询字段
	sqlstr, args := query.BuildQuery()

	//预加载sql
	stmt, err := s.db.Prepare(sqlstr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	//执行sql
	qu, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	// 扫描sql
	for qu.Next() {
		user := rbac.User{}
		qu.Scan(&user.UserId, &user.UserName, &user.Pwd, &user.Email, &user.Phone, &user.Address, &user.State)
		userSet.Items = append(userSet.Items, &user)
	}
	//s.log.Infof("userSet: %s", userSet)

	return userSet, nil
}

func (s *service) DeleteUser(ctx context.Context, req *rbac.DeleteUserRequest) (*rbac.User, error) {
	user := rbac.User{}
	// 删除时先查询
	stmt, err := s.db.Prepare(SelectUserSQL)
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(req.UserId).Scan(&user.UserId, &user.UserName,
		&user.Pwd, &user.Email, &user.Phone, &user.Address, &user.State)
	if err != nil {
		return nil, err
	}

	//req.UserId
	stmt, err = s.db.Prepare(DeleteUserSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	exec, err := stmt.Exec(req.UserId)
	if err != nil {
		return nil, err
	}
	num, err2 := exec.RowsAffected()
	if err2 != nil {
		return nil, err2
	}

	if num == 1 {
		return &user, nil
	} else {
		return nil, fmt.Errorf("删除用户失败")

	}
}
