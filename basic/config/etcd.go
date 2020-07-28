package config

// EtcdConfig Etcd 配置
type IEtcdConfig interface {
	GetEnabled() bool
	GetPort() int
	GetHost() string
}

// defaultEtcdConfig认 etcd 配置
type EtcdConfig struct {
	Enabled bool   `json:"enabled"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
}

func (c *EtcdConfig) GetName() string {
	return "etcd"
}

// GetPort consul 端口
func (c EtcdConfig) GetPort() int {
	return c.Port
}

// GetEnabled consul 激活
func (c EtcdConfig) GetEnabled() bool {
	return c.Enabled
}

// GetHost consul 主机地址
func (c EtcdConfig) GetHost() string {
	return c.Host
}
