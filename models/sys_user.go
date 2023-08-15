package models

import "gorm.io/gorm"

type SysUser struct {
	gorm.Model
	Username  string `gorm:"column:username;type:varchar(50);" json:"username"`
	Password  string `gorm:"column:password;type:varchar(36);" json:"password"`
	Phone     string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	WxUnionId string `gorm:"column:wx_union_id;type:varchar(255);" json:"wx_union_id"`
	WxOpenId  string `gorm:"column:wx_open_id;type:varchar(255);" json:"wx_open_id"`
	Avatar    string `gorm:"column:avatar;type:varchar(255);" json:"avatar"` // 头像
	Remarks   string `gorm:"column:remarks;type:varchar(255);" json:"remarks"`
}

func (user *SysUser) TableName() string {
	return "sys_user"
}

// GetUserByUsernamePassword 根据用户名和密码查询数据
func GetUserByUsernamePassword(username, password string) (*SysUser, error) {
	data := new(SysUser)
	err := DB.Where("username = ? and password = ?", username, password).First(data).Error
	return data, err
}

func GetUserList(keyword string) *gorm.DB {
	tx := DB.Model(new(SysUser)).Select("username,phone,avatar")
	if keyword != "" {
		tx.Where("username LIKE ?", "%"+keyword+"%")
	}
	return tx
}

// GetUserDetail 根据ID获取管理员信息
func GetUserDetail(id uint) (*SysUser, error) {
	su := new(SysUser)
	err := DB.Model(new(SysUser)).Where("id = ?", id).First(su).Error
	return su, err
}
