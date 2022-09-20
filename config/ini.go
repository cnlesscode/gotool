package config

import "gopkg.in/ini.v1"

var Ini *ini.File

func init() {
	iniFile, err := ini.Load("./config.ini")
	if err != nil {
		panic("配置文件 config.ini 加载错误")
	} else {
		Ini = iniFile
	}
}
