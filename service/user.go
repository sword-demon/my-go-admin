package service

import (
	"github.com/gin-gonic/gin"
	"my-go-admin/models"
	"net/http"
	"strconv"
)

// GetUserList 获取管理员列表
func GetUserList(c *gin.Context) {
	in := &GetUserListRequest{NewQueryRequest()}
	// query 查询 绑定query查询参数
	err := c.ShouldBindQuery(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	var (
		cnt  int64
		list = make([]*GetUserListReply, 0)
	)
	err = models.GetUserList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": gin.H{
			"list":  list,
			"count": cnt,
		},
	})
}

// AddUser 新增管理员信息
func AddUser(c *gin.Context) {
	in := new(AddUserRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否存在
	var cnt int64
	err = models.DB.Model(new(models.SysUser)).Where("username = ?", in.Username).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常，保存失败！",
		})
		return
	}

	// 大于0说明用户已经存在
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}

	// 2. 保存用户数据
	err = models.DB.Create(&models.SysUser{
		Username: in.Username,
		Password: in.Password,
		Phone:    in.Phone,
		Remarks:  in.Remarks,
		Email:    in.Email,
	}).Error

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常,保存失败！",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "保存成功",
	})

}

// GetUserDetail 根据ID获取管理员详情信息
func GetUserDetail(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参数不能为空",
		})
		return
	}
	uId, err := strconv.Atoi(id)
	data := new(GetUserDetailReply)
	// 1. 获取角色基本信息
	sysUser, err := models.GetUserDetail(uint(uId))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	data.ID = sysUser.ID
	data.Remarks = sysUser.Remarks
	data.Phone = sysUser.Phone
	data.Username = sysUser.Username
	data.Password = sysUser.Password
	data.Email = sysUser.Email
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    "获取成功",
		"result": data,
	})
}

// UpdateUser 修改管理员信息
func UpdateUser(c *gin.Context) {
	in := new(UpdateUserRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断用户名是否已存在
	var cnt int64
	err = models.DB.Model(new(models.SysUser)).Where("id != ? AND username = ?", in.ID, in.Username).Count(&cnt).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	if cnt > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "用户已存在",
		})
		return
	}

	// 2. 修改数据
	err = models.DB.Model(new(models.SysUser)).Where("id = ?", in.ID).Updates(map[string]any{
		"password": in.Password,
		"username": in.Username,
		"phone":    in.Phone,
		"remarks":  in.Remarks,
	}).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// DeleteUser 删除管理员信息
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "必填参不能为空",
		})
		return
	}

	// 删除管理员
	err := models.DB.Where("id = ?", id).Delete(new(models.SysUser)).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据库异常",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}
