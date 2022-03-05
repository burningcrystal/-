package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

//本文件应用于做一些配置数据的处理

var(
	AppMode string
	HttpPort string
	JwtKey string
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassWord string
	DbName string
	)

func init(){
	file,err := ini.Load("config/config.ini")
	if err!=nil{
		fmt.Println("输入的配置文件路径错误")
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug") //mustring的作用是方式实在取不到值，设置一个备用值
	HttpPort = file.Section("server").Key("HttpPort").MustString("3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("fhgcnfshdgce")
}

func LoadData(file *ini.File) {
	Db  = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("am")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("graff^")
	DbName  = file.Section("database").Key("DbName").MustString("calloflife")
}