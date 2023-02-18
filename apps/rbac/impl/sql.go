package impl

const (
	insertUserSQL = `INSERT INTO s_user (
		 user_name, pwd, email, phone, address, state, user_type
	) VALUES (?,?,?,?,?,?,?);`
	//INSERT INTO `s_user` (user_name,pwd,email,phone,address,state,user_type) VALUES ('wdf','123456','125554566@qq.com','13333728570','河南',0,'管理员');

	QueryUsersSQL = `select user_id,user_name,pwd,email,phone,address,state from s_user`

	SelectUserSQL = `select user_id,user_name,pwd,email,phone,address,state from s_user where user_id = ?;`
	DeleteUserSQL = `delete from s_user where user_id = ?;`
)

const (
	insertRole = `insert into s_role(role_name, description) VALUES (?,?);`
	deleteRole = `delete from s_role where role_id = ?;`
	selectRole = "select role_id, role_name, description from s_role "
)
