package constant

const (
	RESPONSE_CODE_OK     = 0   // 正常响应
	RESPONSE_CODE_ERROR  = 100 // 常规错误
	RESPONSE_CODE_SYSTEM = 200 // 系统故障

	INVALID_PARAMS = 101

	RESPONSE_CODE_SESSION_INVALID   = 300 // 登录会话无效或已掉线
	RESPONSE_CODE_SESSION_REPLACED  = 301 // 登录被顶替
	RESPONSE_CODE_NO_API_PERMISSION = 350 // 无管理接口权限

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)

var MsgFlags = map[int]string{
	RESPONSE_CODE_OK: "OK",
	INVALID_PARAMS:   "请求参数错误",

	RESPONSE_CODE_SESSION_INVALID:   "登录超时，请重新登录",
	RESPONSE_CODE_SESSION_REPLACED:  "账号在其他地方登录，请注意账号安全",
	RESPONSE_CODE_NO_API_PERMISSION: "对不起，你没有权限",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",

	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "保存图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "检查图片失败",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "校验图片错误，图片格式或大小有问题",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return ""
}
