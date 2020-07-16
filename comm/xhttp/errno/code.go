/*
 * @Author       : jianyao
 * @Date         : 2020-07-14 06:29:27
 * @LastEditTime : 2020-07-14 06:46:31
 * @Description  : file content
 */ 
package errno

var (
	OK                  = &Errno{Code: 0, Message: "OK"}
	BadRequestError     = &Errno{Code: 400, Message: "Bad Request"}
	UnauthorizedError   = &Errno{Code: 401, Message: "Unauthorized"}
	ForbiddenError      = &Errno{Code: 403, Message: "Forbidden"}
	InternalServerError = &Errno{Code: 500, Message: "internal server error"}
	InvalidTokenError   = &Errno{Code: 600, Message: "Invalid token"}
)

var (
	ErrUserLogin    =  &Errno{Code: 1001, Message: "login fail"}
	ErrSendSms      =  &Errno{Code: 1002, Message: "send sms fail"}
	ErrUserAddUserOauth =   &Errno{Code: 1002, Message: "db add  UserOauth fail"}
	ErrGetFriendHelp =   &Errno{Code: 1003, Message: "获取好友助力失败"}
	ErrDailyTask =   &Errno{Code: 1004, Message: "获取每日任务失败"}
	ErrSmsCodeInvalid =   &Errno{Code: 1005, Message: "短信验证码校验错误"}
)