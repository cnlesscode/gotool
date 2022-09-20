### Mysql 操作工具
gotool Mysql 操作工具以 gorm <https://gorm.io/zh_CN/docs> 为基础进行了基于连接池的封装，使用更加便利~

#### 配置数据库
请打开( 不存在则创建 ) ./config.ini，添加或修改如下配置 :
```ini
# Mysql 数据库设置
[Mysql]
Host=tcp(127.0.0.1:3306)
User=root
Password=root
ConnMaxLifetime=60
MaxOpenConns=100
MaxIdleConns=50
TableName=lesscode
TablePrefix=t_
Charset=utf8mb4
```
##### 多个数据库配置 
请打开( 不存在则创建 ) ./config.ini，添加或修改如下配置 :
```ini
# Mysql 数据库设置
[Mysql]
Host=tcp(127.0.0.1:3306)
User=root
Password=root
ConnMaxLifetime=60
MaxOpenConns=100
MaxIdleConns=50
TableName=lesscode
TablePrefix=t_
Charset=utf8mb4
[Mysql2]
Host=tcp(192.168.188:3306)
.... 同上配置 ...
[Mysql...]
Host=tcp(192.168.101:3306)
.... 同上配置 ...
```

#### 工具加载
```go
import (
	"github.com/cnlesscode/gotool/db"
)
```

#### db.init() 获取 gorm 操作对象 
功能 : 获取 gorm 操作对象 
参数 : 数据库配置名称，可选参数， 默认 Mysql
说明 : 通过不同配置可以获取基于不同数据库的 orm 对象，实现多库操作
示例 : 
```go
package main

import (
	"fmt"
	"github.com/cnlesscode/gotool/db"
)

func main() {
	db := db.Init()
	fmt.Printf("db: %v\n", db)
}
```

#### gorm 更多用法
获得 gorm 操作对象后您就可以进行任意形式的数据操作，关于 gorm 的知识请访问
gorm 官网 : <https://gorm.io/zh_CN/docs> 
简单示例
```go
package main
import (
	"fmt"
	"github.com/cnlesscode/gotool/db"
)
type Students struct {
	Id      int    `gorm:"column:st_id;primaryKey"`
	ClassId int    `gorm:"column:st_class_id" binding:"gt=0"`
	Name    string `gorm:"column:st_name" binding:"min=2,max=20"`
	Age     int    `gorm:"column:st_age" binding:"gt=5,lt=200"`
}
func main() {
	db := db.Init()
	var sts []Students
	err := db.Limit(10).Find(&sts).Error
	if err == nil {
		fmt.Printf("sts: %v\n", sts)
	} else {
		fmt.Printf("err: %v\n", err)
	}
}
```