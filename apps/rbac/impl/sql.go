package impl

const (
	insertUserSQL = `INSERT INTO s_user (
		 user_name, pwd, email, phone, address, state, user_type
	) VALUES (?,?,?,?,?,?,?);`
	//INSERT INTO `s_user` (user_name,pwd,email,phone,address,state,user_type) VALUES ('wdf','123456','125554566@qq.com','13333728570','河南',0,'管理员');
)
