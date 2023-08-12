// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"gorm.io/gorm"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增id" json:"id"`              // 自增id
	Username   string         `gorm:"column:username;not null;comment:用户名" json:"username"`                        // 用户名
	Password   string         `gorm:"column:password;not null;comment:密码" json:"password"`                         // 密码
	CreateTime int64          `gorm:"column:create_time;type:int;type:unsigned;autoCreateTime" json:"create_time"` // 创建时间
	UpdateTime int64          `gorm:"column:update_time;type:int;type:unsigned;autoUpdateTime" json:"update_time"` // 更新时间
	DeleteTime gorm.DeletedAt `gorm:"column:delete_time;comment:删除时间" json:"delete_time"`                          // 删除时间
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
