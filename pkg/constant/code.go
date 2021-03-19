package constant

const (
	RESPONSE_CODE_OK                = 0   // 正常响应
	RESPONSE_CODE_ERROR             = 100 // 常规错误
	RESPONSE_CODE_SYSTEM            = 200 // 系统故障
	RESPONSE_CODE_SESSION_INVALID   = 300 // 登录会话无效或已掉线
	RESPONSE_CODE_SESSION_REPLACED  = 301 // 登录被顶替
	RESPONSE_CODE_NO_API_PERMISSION = 350 // 无管理接口权限
	RESPONSE_CODE_NO_USER           = 10001
)

func GetResponseMsg(code int) string {
	msg := ""
	switch code {
	case RESPONSE_CODE_SESSION_INVALID:
		msg = "登录超时，请重新登录"
	case RESPONSE_CODE_SESSION_REPLACED:
		msg = "账号在其他地方登录，请注意账号安全"
	case RESPONSE_CODE_NO_API_PERMISSION:
		msg = "对不起，你没有权限"
	}
	return msg
}
