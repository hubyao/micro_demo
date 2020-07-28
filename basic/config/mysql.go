package config

import "time"

// MysqlConfig mysql 配置 接口
type IMysqlConfig interface {
	GetURL() string
	GetEnabled() bool
	GetMaxIdleConnection() int
	GetMaxOpenConnection() int
	GetConnMaxLifetime() time.Duration
}

// defaultMysqlConfig mysql 配置
type MysqlConfig struct {
	URL               string        `json:"url"`
	Enable            bool          `json:"enabled"`
	MaxIdleConnection int           `json:"maxIdleConnection"`
	MaxOpenConnection int           `json:"maxOpenConnection"`
	ConnMaxLifetime   time.Duration `json:"connMaxLifetime"`
}

func (m *MysqlConfig) GetData(data interface{}) interface{} {
	return data.(*MysqlConfig)
}
func (m *MysqlConfig) GetName() string {
	return "mysql"
}

// URL mysql 连接
func (m *MysqlConfig) GetURL() string {
	return m.URL
}

// Enabled 激活
func (m *MysqlConfig) GetEnabled() bool {
	return m.Enable
}

// 闲置连接数
func (m *MysqlConfig) GetMaxIdleConnection() int {
	return m.MaxIdleConnection
}

// 打开连接数
func (m *MysqlConfig) GetMaxOpenConnection() int {
	return m.MaxOpenConnection
}

// 连接数断开时间
func (m *MysqlConfig) GetConnMaxLifetime() time.Duration {
	return m.ConnMaxLifetime
}
