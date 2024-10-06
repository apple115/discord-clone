package models

import (
	"discord-clone/pkg/setting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	log.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}
	// 自动迁移数据库表
	err = db.AutoMigrate(
		&User{},
		&Guild{},
		&UserGuild{},
		&Channel{},
		&UserChannel{},
		&UserActivityLog{},
		&Role{},
		&Permission{},
		&UserRole{},
		&RolePermission{},
	)
	if err != nil {
		log.Fatalf("AutoMigrate err: %v", err)
	}
}
