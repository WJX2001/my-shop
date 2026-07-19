package types

// 错误码定义
const (
	ReturnSuccess             = 2000 // 成功返回
	SystemDbErr               = 3000 // 数据库错误
	InvalidFormatError        = 3001 // 无效的参数格式
	InvalidVerifyWay          = 3002 // 无效的验证方式
	ParamEmptyError           = 3003 // 传入参数为空
	UserToKenCheckError       = 3004 // 用户 Token 校验失败
	PhoneFormatError          = 4003 // 手机号码格式不正确
	PhoneVerifyCodeEmptyError = 4004 // 手机号码验证码为空
	PhoneVerifyCodeError      = 4005 // 手机号码验证码不正确
	EmailFormatError          = 4007 // 邮箱格式不正确
	EmailVerifyCodeEmptyError = 4008 // 邮箱码验证码为空
	EmailVerifyCodeError      = 4009 // 邮箱验证码不正确
	UserAlreadyRegister       = 4010 // 用户已经注册
	UserNotRegister           = 4011 // 用户还没有注册
	NoThisLoginRegisterWay    = 4012 // 没有这种登陆注册验证方式
	UserIsNotExist            = 4013 // 没有这个用户
	UserIsExist               = 4014 // 用户已经存在
	InviteCodeNotExist        = 4015 // 没有这个邀请码
	CreateUserFail            = 4017 // 创建用户失败
	InsertIntegralFail        = 4018 // 插入积分失败
	CreateUserWalletFail      = 4019 // 创建用户钱包失败
	PasswordError             = 4020 // 输入的密码错误
	PasswordIsEmpty           = 4021 // 输入密码为空
	NewOldPasswordEqual       = 4022 // 新旧密码相等
	TwicePasswordNotEqual     = 4023 // 新旧密码相等
	AddLoginTimesError        = 4024 // 添加登陆次数错误
	AlreadyBindPassword       = 4052 // 已经绑定支付密码
)
