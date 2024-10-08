package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

type MongoDB struct {
	URI      string
	Database string
}

var MongoDBSetting = &MongoDB{}

type MongoDB struct {
	URI      string
	Database string
}

var MongoDBSetting = &MongoDB{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup,fail to parse 'conf/app.ini':%v", err)
	}
	mapTo("database", DatabaseSetting)
	mapTo("server", ServerSetting)
	mapTo("mongodb", MongoDBSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err:%v", section, err)
	}
}
