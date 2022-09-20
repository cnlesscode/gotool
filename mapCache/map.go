package mapCache

import (
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cnlesscode/gotool/config"
	"github.com/cnlesscode/gotool/gstring"
	"github.com/cnlesscode/gotool/slice"
)

// 缓存数据读取函数
type CacheFunc func(args ...any) (any, error)

// 组合缓存名称
func InitCacheName(keyName string, cacheParameters ...any) string {
	name := keyName + slice.ToString(cacheParameters, "")
	name = gstring.StripPunctuation(name)
	return name
}

var MapCacher sync.Map
var MapCacherExpiration sync.Map

// 初始化
func init() {
	go (func() {
		// 缓存有效期检查
		GCIntervalTime, errIni := config.Ini.Section("MapCache").Key("GCIntervalTime").Int64()
		if errIni != nil {
			GCIntervalTime = 600
		} else {
			if GCIntervalTime < 10 {
				GCIntervalTime = 600
			}
		}
		println("✔ 缓存有效期检查线程已经启动，间隔 : " + strconv.Itoa(int(GCIntervalTime)) + " 秒\n")
		for {
			time.Sleep(time.Second * time.Duration(GCIntervalTime))
			MapCacherExpiration.Range(func(key, value any) bool {
				// 比对有效期
				now := time.Now().Unix()
				expTime := value.(int64)
				if now > expTime {
					MapCacherExpiration.Delete(key)
					MapCacher.Delete(key)
				}
				return true
			})
		}
	})()
}

// 获取变量 不存在则自动设置
func Cache(keyName string, expiration int, cacheFunc CacheFunc, cacheParameters ...any) any {
	var data any
	var ok bool
	var err error
	keyName = InitCacheName(keyName, cacheParameters...)
	data, ok = MapCacher.Load(keyName)
	if ok {
		// 有效期检查
		if expiration > 1 {
			expForCurrent, expOk := MapCacherExpiration.Load(keyName)
			if expOk {
				now := time.Now().Unix()
				// 过期缓存
				if now > expForCurrent.(int64) {
					MapCacherExpiration.Delete(keyName)
					MapCacher.Delete(keyName)
					// 重新设置缓存
					data, err = cacheFunc(cacheParameters...)
					if err == nil {
						Set(keyName, expiration, data)
					}
					return data
				}
			}
		}
		return data
	}
	data, err = cacheFunc(cacheParameters...)
	if err == nil {
		Set(keyName, expiration, data)
	}
	return data
}

// 设置缓存
func Set(keyName string, expiration int, data any) {
	MapCacher.Store(keyName, data)
	if expiration > 0 {
		MapCacherExpiration.Store(keyName, time.Now().Unix()+int64(expiration))
	}
}

// 删除 模糊搜索方式
func Remove(name string) {
	if name == "" {
		return
	}
	MapCacher.Range(func(key, value any) bool {
		if strings.Contains(key.(string), name) {
			MapCacher.Delete(key)
		}
		return true
	})
}

// 清空
func Clear() {
	MapCacher = sync.Map{}
	MapCacherExpiration = sync.Map{}
}
