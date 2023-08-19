package service

import (
	"github.com/gin-gonic/gin"
	"my-go-admin/models"
	"net/http"
)

// GetRoleList 角色列表
func GetRoleList(c *gin.Context) {
	in := &GetRoleListRequest{NewQueryRequest()}
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
		list = make([]*GetRoleListReply, 0)
	)
	err = models.GetRoleList(in.Keyword).Count(&cnt).Offset((in.Page - 1) * in.Size).Limit(in.Size).Find(&list).Error
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

// AddRole 新增角色信息
func AddRole(c *gin.Context) {
	in := new(AddRoleRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 判断角色名称是否存在
	var cnt int64
	err = models.DB.Model(new(models.SysRole)).Where("name = ?", in.Name).Count(&cnt).Error
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
	err = models.DB.Create(&models.SysRole{
		Name:    in.Name,
		IsAdmin: in.IsAdmin,
		Sort:    in.Sort,
		Remarks: in.Remarks,
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
