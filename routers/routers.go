// 路由层 管理程序的路由信息
package routers

import (
	"golang_project_scaffold/routers/auth"

	"github.com/gin-gonic/gin"
)

// 注册路由的方法
func RegistrerRouters(r *gin.Engine) {
	// 登录的路由配置
	// 1. 登录: login
	// 2. 退出: logout
	// 3. /api/auth/login
	// 4. /api/auth/logout
	apiGroup := r.Group("/api")
	auth.RegisterSubRouter(apiGroup)
}
