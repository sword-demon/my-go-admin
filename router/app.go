package router

import (
	"github.com/gin-gonic/gin"
	"my-go-admin/middleware"
	"my-go-admin/service"
)

func App() *gin.Engine {
	r := gin.Default()
	// 添加跨域中间件
	r.Use(middleware.Cors())

	// 根据用户名和密码登录
	r.POST("/login/password", service.LoginPassword)

	// 管理员 start
	// 管理员列表
	r.GET("/user", service.GetUserList)
	// 新增管理员
	r.POST("/user", service.AddUser)
	// 管理员详情
	r.GET("/user/detail", service.GetUserDetail)
	// 修改管理员
	r.PUT("/user", service.UpdateUser)
	// 删除管理员
	r.DELETE("/user/:id", service.DeleteUser)
	// 管理员 end

	// 角色管理 start
	// 角色列表
	r.GET("/role", service.GetRoleList)
	// 新增角色
	r.POST("/role", service.AddRole)
	r.GET("/role/detail", service.GetRoleDetail)
	// 修改角色
	r.PUT("/role", service.UpdateRole)
	// 删除角色
	r.DELETE("/role/:id", service.DeleteRole)
	// 角色管理 end

	return r
}
