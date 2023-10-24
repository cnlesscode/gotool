package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client        *redis.Client
	Options       redis.Options
	VarNamePrefix string
}

func Init(Addr string, Password string, DB int, varNamePrefix string) (m *Redis) {
	return &Redis{
		Options: redis.Options{
			Addr:     Addr,
			Password: Password,
			DB:       DB,
		},
		VarNamePrefix: varNamePrefix,
	}
}

// 初始化 redis 客户端
func (m *Redis) InitRedisClient() {
	// 初始化 redis 客户端
	m.Client = redis.NewClient(&m.Options)
}

// 设置变量
func (m *Redis) Set(name string, value interface{}, expiration int64) error {
	if m.Client == nil {
		m.InitRedisClient()
	}
	return m.Client.Set(
		context.Background(),
		m.VarNamePrefix+name,
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
		m.VarNamePrefix+name,
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
func (m *Redis) DoGet(args ...any) (any, error) {
	if m.Client == nil {
		m.InitRedisClient()
	}
	cmd := m.Client.Do(
		context.Background(),
		args...,
	)
	return cmd.Result()
}

func (m *Redis) Close() error {
	if m.Client == nil {
		return nil
	}
	return m.Client.Close()
}
