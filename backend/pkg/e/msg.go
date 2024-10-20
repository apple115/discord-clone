package e

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_NOT_EXIST_CHANNEL:        "频道不存在",
	ERROR_CHECK_EXIST_CHANNEL_FAIL: "确认频道存在失败",
	ERROR_ADD_CHANNEL_FAIL:         "频道添加失败",
	ERROR_DELETE_CHANNEL_FAIL:      "频道删除失败",
	ERROR_EDIT_CHANNEL_FAIL:        "频道编辑失败",
	ERROR_GET_CHANNEL_FAIL:         "获取频道失败",
	ERROR_GET_CHANNEL_BY_ID_FAIL:   "通过ID获取频道失败",

	ERROR_GET_CHANNEL_MESSAGE_FAIL: "获取频道历史消息失败",

	ERROR_HASHPASSWORD:      "密码hash失败",
	ERROR_EXIST_USER:        "用户已存在",
	ERROR_EXIST_EMAIL:       "邮箱已存在",
	ERROR_CHECK_EXIST_USER:  "确认用户存在失败",
	ERROR_CHECK_EXIST_EMAIL: "确认邮箱存在失败",
	ERROR_ADD_USER:          "添加用户失败",
	ERROR_GET_USER:          "获取用户失败",
	ERROR_NOT_EXIST_USER: "用户不存在",

	ERROR_GEN_TOKEN:                "生成token失败",
	ERROR_VERIFY_RFRESH_TOKEN:      "验证refresh token失败",
	ERROR_VERIFY_ACCESS_TOKEN:      "验证access token失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "验证token超时",

	ERROR_CAPTCHA_KEY_EMPTY:  "验证码key为空",
	ERROR_CAPTCHA_VERIFYDOTS: "验证码验证失败",
	ERROR_CAPTCHA_STORE:      "非法存储格式",

	ERROR_UPLOAD_CHECK_IMAGE_FAIL:   "上传图片验证错误",
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT: "上传图片格式错误",
	ERROR_UPLOAD_SAVE_IMAGE_FAIL:    "上传图片保存错误",

	FAIL_CAPTCHA: "验证码错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
