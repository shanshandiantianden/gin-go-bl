package errmsg

// 错误码定义规则
// 1. 服务级别为 1 2 3......
// 2. 模块级别为 01 02 03......
// 3. 具体错误码为 01 02 03......
// 4. 错误信息为 错误码 + 错误信息
// 5. 错误码为 0 表示成功
// 6. 错误码为 1 表示失败
// 7. 示例错误码为 1 00 01 表示服务异常，请联系管理员

var (
	// OK
	OK = NewError(0, "OK")

	// 服务级错误码
	ErrServer    = NewError(10001, "服务异常")
	ErrParam     = NewError(10002, "参数有误")
	ErrSignParam = NewError(10003, "签名参数有误")
	ErrGormQuery = NewError(10004, "Gorm执行错误")
	// 模块级错误码 - 用户模块
	ErrUserPhone       = NewError(20101, "用户手机号不合法")
	ErrUserCaptcha     = NewError(20102, "用户验证码有误")
	ErrUserName        = NewError(20103, "用户名不合法")
	ErrUserPassword    = NewError(20104, "密码错误")
	ErrLoginNil        = NewError(20105, "用户名或密码为空")
	ErrUserNotExist    = NewError(20106, "用户不存在")
	ErrUserExist       = NewError(20107, "用户已存在")
	ErrUserToken       = NewError(20108, "用户token有误")
	ErrUserPwd         = NewError(20109, "用户密码有误")
	ErrUnauthorized    = NewError(20110, "用户未授权")
	ErrArticleNotExist = NewError(20106, "文章为空")

	// ...
)
