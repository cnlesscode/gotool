package gotool

import (
	"fmt"
	"testing"

	"github.com/cnlesscode/gotool/db"
)

// 单元测试

// 测试命令 : go test -v -run=TestDB
func TestDB(t *testing.T) {
	dbConfigs := map[string]map[string]string{
		// "DB": {
		// 	"DBType":       "MSSQL",
		// 	"RunMode":      "dev",
		// 	"HostDev":      "localhost",
		// 	"UsernameDev":  "root",
		// 	"PasswordDev":  "root",
		// 	"Port":         "1433",
		// 	"DBName":       "test",
		// 	"Charset":      "utf8mb4",
		// 	"MaxOpenConns": "1000",
		// 	"MaxIdleConns": "200",
		// 	"MaxLifetime":  "3600",
		// },
		"DB": {
			"DBType":       "MySQL",
			"RunMode":      "dev",
			"HostDev":      "localhost",
			"UsernameDev":  "root",
			"PasswordDev":  "root",
			"Port":         "3306",
			"DBName":       "test",
			"Charset":      "utf8mb4",
			"MaxOpenConns": "1000",
			"MaxIdleConns": "200",
			"MaxLifetime":  "3600",
		},
	}
	db.Start(dbConfigs)
	gormDB := db.Init()
	fmt.Printf("gormDB: %v\n", gormDB)
}
