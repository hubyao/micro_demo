package config

// Global 配置
type Global struct {
	Env string `json:"env"` // 环境变量
}

// IsEnvLoc 是否是本地环境
//func (g *Global) IsEnvLoc() bool {
//	return g.Env == common.EnvLoc
//}
//
//// IsEnvDev 是否是开发环境:true=是,false=不是
//func (g *Global) IsEnvDev() bool {
//	return g.Env == common.EnvDev
//}
//
//// IsEnvProd 是否是正式环境
//func (g *Global) IsEnvProd() bool {
//	return g.Env == common.EnvProd
//}
