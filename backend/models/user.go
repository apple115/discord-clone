package models

import (
	"gorm.io/gorm"
)

// User 结构体代表用户表
type User struct {
	gorm.Model
	Username          string `gorm:"type:varchar(100);unique_index"`
	PasswordHash      string `gorm:"type:varchar(100)"`
	Email             string `gorm:"type:varchar(100);unique_index"`
	ProfilePictureUrl string `gorm:"type:varchar(100)"`
	StatueMessage     string `gorm:"type:enum('在线', '注销','隐身');not null"`
}

// Guild 结构体代表群组表
type Guild struct {
	gorm.Model
	GuildName   string `gorm:"type:varchar(255);not null"`
	OwnerID     uint   // 虽然数据库中是 int，但这里使用 uint 以便和 gorm.Model 的 ID 类型一致
	Description string
}

// UserGuild 结构体代表用户与群组关系表
type UserGuild struct {
	gorm.Model
	UserID   uint // 虽然数据库中是 int，但这里使用 uint 以便和 gorm.Model 的 ID 类型一致
	GuildID  uint
	JoinedAt gorm.DeletedAt // 使用 gorm.DeletedAt 类型可以方便地表示加入时间，也可以在软删除场景下使用
}

// Channel 结构体代表频道表
type Channel struct {
	gorm.Model
	ChannelName string `gorm:"type:varchar(255);not null"`
	GuildID     uint   // 虽然数据库中是 int，但这里使用 uint 以便和 gorm.Model 的 ID 类型一致
	Description string
}

// UserChannel 结构体代表用户与频道关系表
type UserChannel struct {
	gorm.Model
	UserID    uint
	ChannelID uint
	IsMuted   bool `gorm:"default:false"`
}

// UserActivityLog 结构体代表用户活动日志表
type UserActivityLog struct {
	gorm.Model
	UserID    uint   // 虽然数据库中是 int，但这里使用 uint 以便和 gorm.Model 的 ID 类型一致
	Action    string `gorm:"type:varchar(255);not null"`
	IPAddress string `gorm:"type:varchar(45)"`
}

// Role 结构体代表角色表
type Role struct {
	gorm.Model
	RoleName    string `gorm:"type:varchar(255);not null"`
	Description string
}

// Permission 结构体代表权限表
type Permission struct {
	gorm.Model
	PermissionName string `gorm:"type:varchar(255);not null"`
	Description    string
}

// UserRole 结构体代表用户角色关系表
type UserRole struct {
	gorm.Model
	UserID uint
	RoleID uint
}

// RolePermission 结构体代表角色权限关系表
type RolePermission struct {
	gorm.Model
	RoleID       uint
	PermissionID uint
}