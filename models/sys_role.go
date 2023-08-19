package models

import (
	"gorm.io/gorm"
)

type SysRole struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(30)" json:"name"`
	// 是否是超管 【0-否 1-是】
	IsAdmin int8   `gorm:"column:is_admin;type:tinyint(1);default:0" json:"is_admin"`
	Sort    int64  `gorm:"column:sort;type:int(11);default:1" json:"sort"`
	Remarks string `gorm:"remarks;type:varchar(255);" json:"remarks"`
}

func (role *SysRole) TableName() string {
	return "sys_role"
}

// GetRoleList 获取角色列表
func GetRoleList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysRole)).Select("id, name, is_admin, sort, created_at, updated_at")
	if keyword != "" {
		tx.Where("name LIKE ?", "%"+keyword+"%")
	}
	tx.Order("sort ASC")
	return tx
}
