package all

import (
	// 注册所有HTTP服务模块, 暴露给框架HTTP服务器加载
	_ "github.com/defeng-hub/mcube-demo/apps/book/api"
	_ "github.com/defeng-hub/mcube-demo/apps/rbac/api"
)
