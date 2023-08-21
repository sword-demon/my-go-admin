package service

import (
	"github.com/gin-gonic/gin"
	"my-go-admin/models"
	"net/http"
)

// GetMenuList 设置-获取菜单列表
func GetMenuList(c *gin.Context) {
	Menus(c)
}

// Menus 获取菜单列表
func Menus(c *gin.Context) {
	data := make([]*MenuReply, 0)
	allMenus := make([]*AllMenu, 0)

	// 获取所有菜单列表数据
	tx := models.GetMenusList()

	err := tx.Find(&allMenus).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "数据异常",
		})
	}

	data = allMenuToMenuReply(allMenus)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "加载成功",
		"data": data,
	})
}

// 生成树形菜单
func allMenuToMenuReply(allMenu []*AllMenu) []*MenuReply {
	reply := make([]*MenuReply, 0)
	// 一层循环，得到顶层菜单
	for _, v := range allMenu {
		if v.ParentId == 0 {
			reply = append(reply, &MenuReply{
				ID:       v.ID,
				Name:     v.Name,
				WebIcon:  v.WebIcon,
				Sort:     v.Sort,
				Path:     v.Path,
				Level:    v.Level,
				SubMenus: getChildrenMenu(v.ID, allMenu),
			})
		}
	}
	return reply
}

// getChildrenMenu 获取子菜单
func getChildrenMenu(parentId int, allMenus []*AllMenu) []*MenuReply {
	data := make([]*MenuReply, 0)
	for _, v := range allMenus {
		if v.ParentId == parentId {
			data = append(data, &MenuReply{
				ID:       v.ID,
				Name:     v.Name,
				WebIcon:  v.WebIcon,
				Sort:     v.Sort,
				Path:     v.Path,
				Level:    v.Level,
				ParentId: v.ParentId,
				SubMenus: getChildrenMenu(v.ID, allMenus),
			})
		}
	}
	return data
}

// AddMenu 新增菜单
func AddMenu(c *gin.Context) {
	in := new(AddMenuRequest)
	err := c.ShouldBindJSON(in)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数异常",
		})
		return
	}

	// 1. 保存数据
	err = models.DB.Create(&models.SysMenu{
		ParentId: in.ParentId,
		Name:     in.Name,
		WebIcon:  in.WebIcon,
		Sort:     in.Sort,
		Path:     in.Path,
		Level:    in.Level,
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
		"msg":  "新增成功",
	})
}
