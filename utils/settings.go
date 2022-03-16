package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string

	AccessKey  string
	SecretKey  string
	Bucket     string
	QiniuSever string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败, err: ", err)
	}

	loadServer(cfg)
	loadData(cfg)
	loadQiNiu(cfg)

}

func loadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").String()
	HttpPort = cfg.Section("server").Key("HttpPort").String()
	JwtKey = cfg.Section("server").Key("JwtKey").String()
}

func loadData(cfg *ini.File) {
	Db = cfg.Section("database").Key("Db").String()
	DbHost = cfg.Section("database").Key("DbHost").String()
	DbPort = cfg.Section("database").Key("DbPort").String()
	DbUser = cfg.Section("database").Key("DbUser").String()
	DbPassword = cfg.Section("database").Key("DbPassword").String()
	DbName = cfg.Section("database").Key("DbName").String()
}

func loadQiNiu(cfg *ini.File) {
	AccessKey = cfg.Section("qiniu").Key("AccessKey").String()
	SecretKey = cfg.Section("qiniu").Key("SecretKey").String()
	Bucket = cfg.Section("qiniu").Key("Bucket").String()
	QiniuSever = cfg.Section("qiniu").Key("QiniuSever").String()
}
