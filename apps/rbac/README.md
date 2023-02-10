# RBAC标准完整权限管理系统的实现
用户、角色、权限三大核心表，加上用户角色、角色权限两个映射表(用于给用户表关联上权限表)。这样就可以通过登录的用户来获取权限列表，或判断是否拥有某个权限。

# 五张表逻辑关联
１、用户表（s_user）：user_id,user_name、pwd,state,user_type

２、角色表（s_role）：role_id,role_name,description

３、菜单表（s_menu）：menu_id,menu_pid,sort,menu_name,func_name,is_enabled,description

４、用户角色表（s_user_role）：urid、user_id、role_id

５、角色菜单表（s_role_menu）：rmid、role_id、menu_id
 
---


```mysql

-- １、用户表（s_user）：user_id,user_name、pwd,state,user_type
create table s_user
(

    user_id int not null AUTO_INCREMENT,  -- 用户ID
    user_name varchar(100),        -- 用户名
    pwd varchar(50),               -- 密码(md5加密)
    email varchar(100),            -- 邮箱
    phone varchar(11),             -- 手机号
    address varchar(100),          -- 地址
    state int,                     -- 用户状态(0:启用，1禁用）
    user_type varchar(50),         -- 用户类型（用于用户分组，比如管理员，省级，市级，县级,etc.)
    CONSTRAINT pk_s_user_user_id PRIMARY KEY(user_id)  -- 主键
);


-- ２、角色表（s_role）：role_id,role_name,description
create table s_role
(
    role_id int not null AUTO_INCREMENT,  -- 角色ID
    role_name varchar(100),        -- 角色名
    description varchar(1000),     -- 角色描述
    CONSTRAINT pk_s_role_role_id PRIMARY KEY(role_id)  -- 主键
);

-- ３、菜单表（s_menu）：menu_id,menu_pid,sort,menu_name,func_name,is_enabled,description
create table s_menu
(
    menu_id int not null AUTO_INCREMENT,   -- 菜单ID
    menu_name varchar(100),         -- 菜单名
    menu_pid int not null,          -- 菜单父ID （这样可以做到无限级子菜单）
    sort int,                       -- 菜单排序(int型，控制菜单显示顺序）
    func_name varchar(100),         -- 该菜单执行方法函数（可以是url接口地址或者FunctionName）
    is_enabled int,                 -- 是否启用该菜单(int型，控制菜单显示与否,0：启用，1：禁用）
    description varchar(1000),      -- 菜单功能描述
    CONSTRAINT pk_s_menu_menu_id PRIMARY KEY(menu_id)  -- 主键
);

-- ４、用户角色表（s_user_role）：urid、user_id、role_id
create table s_user_role
(
    urid int not null AUTO_INCREMENT,   -- 用户角色ID
    user_id int,         -- 用户ID
    role_id int,         -- 角色ID
    CONSTRAINT pk_s_user_role_urid PRIMARY KEY(urid)  -- 主键
);

-- ５、角色菜单表（s_role_menu）：rmid、role_id、menu_id
create table s_role_menu
(
    rmid int not null auto_increment,   -- 角色菜单ID
    role_id int,         -- 角色ID
    menu_id int,         -- 菜单ID
    CONSTRAINT pk_s_role_menu_rmid PRIMARY KEY(rmid)  -- 主键
);


```


# 操作
```cmd
> protoc -I="."  --go_out=. --go_opt=module="github.com/defeng-hub/mcube-demo" --go-grpc_out=. --go-grpc_opt=module="github.com/defeng-hub/mcube-demo" apps/rbac/pb/*.proto

> protoc-go-inject-tag -input=apps/rbac/*.pb.go

> mcube generate enum -p -m apps/rbac/*.pb.go

```