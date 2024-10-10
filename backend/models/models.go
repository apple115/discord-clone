package models

import (
	"context"
	"discord-clone/pkg/setting"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var MongoClient *mongo.Client
var MongoDB *mongo.Database

func Setup() {
	var err error
	// Gorm 初始化
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
		&ChannelData{},
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

	// MongoDB 初始化
	clientOptions := options.Client().ApplyURI(setting.MongoDBSetting.URI)
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("无法连接 MongoDB: %v", err)
	}
	err = MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("无法 ping MongoDB: %v", err)
	}
	MongoDB = MongoClient.Database(setting.MongoDBSetting.Database)
}
