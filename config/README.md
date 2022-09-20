### 配置工具
gotool 提供基于 .ini 文件的配置工具，基于 gopkg.in/ini.v1


#### 配置文件 ./config.ini
请在项目根目录下创建 ./config.ini 并添加自己的配置， 如 :
```ini
# 运行模式
RunMode=dev

[Cache]
# 支持 redis 和 map
Type=map
Prefix=gotool_

# Redis 缓存配置
[RedisCache]
Host=127.0.0.1
Port=6379
Password=
DefaultDB=0
```
#### 获取配置数据
```go
package main

import (
	"fmt"

	"github.com/cnlesscode/gotool/config"
)

func main() {
	println("welcome to use gotool ...")
	host := config.Ini.Section("RedisCache").Key("Host")
	fmt.Printf("host: %v\n", host)
	port, _ := config.Ini.Section("RedisCache").Key("Port").Int()
	fmt.Printf("host: %v\n", port)
}
```