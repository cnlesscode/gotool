package db

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GoToolDBMap = make(map[string]*gorm.DB)

// 初始化数据库连接池
func Start(dbConfigs map[string]map[string]string) {
	// 遍历数据库配置初始化连接池
	for k, conf := range dbConfigs {
		// 日志级别
		var loggerType logger.LogLevel
		if conf["RunMode"] == "dev" {
			loggerType = logger.Info
		} else {
			loggerType = logger.Error
		}
		// 配置文件
		options := &gorm.Config{
			// 开启事务保证数据一致性
			SkipDefaultTransaction: false,
			// 日志
			Logger: logger.Default.LogMode(loggerType),
			// 命名策略
			NamingStrategy: schema.NamingStrategy{
				// 表前缀
				TablePrefix: conf["TablePrefix"],
				// 单数表名称
				SingularTable: true,
			},
			// 建表时候是否忽略外键
			DisableForeignKeyConstraintWhenMigrating: true,
		}
		// 创建连接池
		var DSN string = ""
		var err error
		if conf["DBType"] == "MySQL" {
			DSN = conf["Username"] + ":" +
				conf["Password"] + "@" +
				"tcp(" + conf["Host"] + ":" + conf["Port"] + ")/" +
				conf["DatabaseName"] + "?charset=" +
				conf["Charset"] + "&parseTime=True&loc=Local"
			GoToolDBMap[k], err = gorm.Open(mysql.Open(DSN), options)
			if err != nil {
				println("✘ 连接库连接池 : " + k + " 初始化失败 ( " + err.Error() + " )")
			} else {
				println("✔ 连接库连接池 : " + k + " 初始化成功 ")
			}
		}
		// 获取基础数据库操作接口
		sqlDB, _ := GoToolDBMap[k].DB()
		//设置数据库连接池参数
		// 最大连接数
		MaxOpenConns, _ := strconv.Atoi(conf["MaxOpenConns"])
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		// 最大空闲连接数
		MaxIdleConns, _ := strconv.Atoi(conf["MaxIdleConns"])
		sqlDB.SetMaxIdleConns(MaxIdleConns)
		// 最大连接时间
		MaxLifetime, _ := strconv.Atoi(conf["MaxLifetime"])
		sqlDB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	}
}

// 获取数据库操作对象
func Init(configName ...string) *gorm.DB {
	if len(configName) < 1 {
		configName = append(configName, "DB")
	}
	gormDB, ok := GoToolDBMap[configName[0]]
	if ok {
		return gormDB
	}
	panic("✘ 数据库连接池 [ " + configName[0] + " ] 初始化失败")
}

// 将 map 对象转换为 sql 条件
func MapToWhere(mapData map[string][]any) (string, []any) {
	var whereSql = make([]string, 0)
	var whereVal = make([]any, 0)
	for k, item := range mapData {
		whereSql = append(whereSql, fmt.Sprintf("%v %v %v ?", item[0], k, item[1]))
		whereVal = append(whereVal, item[2])
	}
	return strings.Join(whereSql, " "), whereVal
}
