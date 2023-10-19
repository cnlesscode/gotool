package iniReader

import "gopkg.in/ini.v1"

type Reader struct {
	iniFile *ini.File
}

func New(configFile string) *Reader {
	iniFile, err := ini.Load(configFile)
	if err != nil {
		panic("配置文件 config.ini 加载错误")
	}
	return &Reader{iniFile: iniFile}
}

// 字符串配置读取
func (m *Reader) String(section, key string) string {
	return m.iniFile.Section(section).Key(key).String()
}

// int 配置读取
func (m *Reader) Int(section, key string) int {
	val := 0
	val, _ = m.iniFile.Section(section).Key(key).Int()
	return val
}

// int64 配置读取
func (m *Reader) Int64(section, key string) int64 {
	var val int64 = 0
	val, _ = m.iniFile.Section(section).Key(key).Int64()
	return val
}

// int32 配置读取
func (m *Reader) Int32(section, key string) int32 {
	val := m.Int64(section, key)
	return int32(val)
}

// float64 配置读取
func (m *Reader) Float64(section, key string) float64 {
	var val float64 = 0
	val, _ = m.iniFile.Section(section).Key(key).Float64()
	return val
}

// float32 配置读取
func (m *Reader) Float32(section, key string) float32 {
	val := m.Float64(section, key)
	return float32(val)
}
