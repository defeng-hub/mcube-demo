CREATE TABLE IF NOT EXISTS `books` (
  `id` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对象Id',
  `create_at` bigint NOT NULL COMMENT '创建时间(13位时间戳)',
  `create_by` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `update_at` bigint NOT NULL COMMENT '更新时间',
  `update_by` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '更新人',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '书名',
  `author` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '作者',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`) USING BTREE COMMENT '用于书名搜索',
  KEY `idx_author` (`author`) USING BTREE COMMENT '用于作者搜索'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


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
    id int not null AUTO_INCREMENT,   -- 用户角色ID
    user_id int,         -- 用户ID
    role_id int,         -- 角色ID
    CONSTRAINT pk_s_user_role_urid PRIMARY KEY(id),  -- 主键
    FOREIGN KEY (user_id) REFERENCES s_user(user_id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES s_user(user_id) ON UPDATE CASCADE,

    FOREIGN KEY (role_id) REFERENCES s_role(role_id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES s_role(role_id) ON UPDATE CASCADE
);

-- ５、角色菜单表（s_role_menu）：rmid、role_id、menu_id
create table s_role_menu
(
    id int not null auto_increment,   -- 角色菜单ID
    role_id int,         -- 角色ID
    menu_id int,         -- 菜单ID
    CONSTRAINT pk_s_role_menu_rmid PRIMARY KEY(rmid)  -- 主键
);
