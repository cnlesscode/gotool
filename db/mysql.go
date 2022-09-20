package db

import (
	"time"

	"github.com/cnlesscode/gotool/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GoToolDBMap = make(map[string]*gorm.DB)

func Init(configName ...string) *gorm.DB {
	if len(configName) < 1 {
		configName = append(configName, "Mysql")
	}
	gormDB, ok := GoToolDBMap[configName[0]]
	if ok {
		return gormDB
	}
	// 日志级别
	var loggerType logger.LogLevel
	if config.Ini.Section("").Key("RunMode").String() == "dev" {
		loggerType = logger.Info
	} else {
		loggerType = logger.Error
	}
	// Mysql 配置
	keys := config.Ini.Section(configName[0]).Keys()
	if len(keys) < 1 {
		panic("✘ 数据库配置错误，请仔细检查")
	}
	Dsn := config.Ini.Section(configName[0]).Key("User").String() + ":" +
		config.Ini.Section(configName[0]).Key("Password").String() + "@" +
		config.Ini.Section(configName[0]).Key("Host").String() + "/" +
		config.Ini.Section(configName[0]).Key("TableName").String() + "?charset=" +
		config.Ini.Section(configName[0]).Key("Charset").String()
	MaxOpenConns, _ := config.Ini.Section(configName[0]).Key("MaxOpenConns").Int()
	MaxIdleConns, _ := config.Ini.Section(configName[0]).Key("MaxIdleConns").Int()

	// 连接 MYSQL
	var err error
	GoToolDBMap[configName[0]], err = gorm.Open(mysql.Open(Dsn), &gorm.Config{
		// 开启事务保证数据一致性
		SkipDefaultTransaction: false,
		// 日志
		Logger: logger.Default.LogMode(loggerType),
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: config.Ini.Section(configName[0]).Key("TablePrefix").String(),
			// 单数表名称
			SingularTable: true,
		},
		// 建表时候是否忽略外键
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("✘ 连接数据库失败, error=" + err.Error())
	} else {
		println("✔ 数据库 " + configName[0] + " 连接成功 ")
	}

	// 获取基础数据库操作接口
	sqlDB, _ := GoToolDBMap[configName[0]].DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(MaxOpenConns)
	sqlDB.SetMaxIdleConns(MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return GoToolDBMap[configName[0]]
}
