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

	return r
}
