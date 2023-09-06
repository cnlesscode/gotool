package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

var redisOptions = &redis.Options{
	Addr:     "",
	Password: "", // 密码
	DB:       0,  // 使用的库
}

var RedisVarNamePrefix = ""

func InitRedis(Addr string, Password string, DB int, VarNamePrefix string) {
	redisOptions.Addr = Addr
	redisOptions.Password = Password
	redisOptions.DB = DB
	RedisVarNamePrefix = VarNamePrefix
}

// 初始化 redis 客户端
func (m *Redis) InitRedisClient() {
	// 初始化 redis 客户端
	m.Client = redis.NewClient(redisOptions)
}

// 设置变量
func (m *Redis) Set(name string, value interface{}, expiration int64) error {
	if m.Client == nil {
		m.InitRedisClient()
	}
	return m.Client.Set(
		context.Background(),
		RedisVarNamePrefix+name,
		value,
		time.Second*time.Duration(expiration),
	).Err()
}

// 获取变量
func (m *Redis) Get(name string) (string, error) {
	if m.Client == nil {
		m.InitRedisClient()
	}
	return m.Client.Get(
		context.Background(),
		RedisVarNamePrefix+name,
	).Result()
}

// 设置变量 Do 模式
func (m *Redis) DoSet(args ...any) error {
	if m.Client == nil {
		m.InitRedisClient()
	}
	cmd := m.Client.Do(
		context.Background(),
		args...,
	)
	return cmd.Err()
}

// 获取变量 Do 模式
func (m *Redis) DoGet(args ...any) ([]string, error) {
	if m.Client == nil {
		m.InitRedisClient()
	}
	cmd := m.Client.Do(
		context.Background(),
		args...,
	)
	return cmd.StringSlice()
}

func (m *Redis) Colse() error {
	if m.Client == nil {
		return nil
	}
	return m.Client.Close()
}
