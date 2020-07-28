package config

import "micro_demo/basic/common"

// defaultProfiles 属性配置文件
type ProfilesConfig struct {
	Env     string `json:"env"`      // 环境变量
	AppName string `json:"app_name"` // 应用名
}

func (p *ProfilesConfig) GetName() string {
	return "profiles"
}

// IsEnvLoc 是否是本地环境
func (p *ProfilesConfig) IsEnvLoc() bool {
	return p.Env == common.EnvLoc
}

// IsEnvDev 是否是开发环境:true=是,false=不是
func (p *ProfilesConfig) IsEnvDev() bool {
	return p.Env == common.EnvDev
}

// IsEnvProd 是否是正式环境
func (p *ProfilesConfig) IsEnvProd() bool {
	return p.Env == common.EnvProd
}

//
func (p *ProfilesConfig) GetAppName() string {
	return p.AppName
}
