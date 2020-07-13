package common

const (
	InventoryHistoryStateNotOut = 1
	InventoryHistoryStateOut    = 2
	RememberMeCookieName        = "remember-me-token"
	TopicPaymentDone            = "mu.micro.book.topic.payment.done"
)

// 环境变量
const (
	EnvLoc = "loc"   // 本地环境
	EnvDev = "dev"   // 开发环境
	EnvProd = "prod" // 正式环境
)
