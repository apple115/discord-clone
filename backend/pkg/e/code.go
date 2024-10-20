package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400 //请求参数错误

	ERROR_NOT_EXIST_CHANNEL        = 10001
	ERROR_CHECK_EXIST_CHANNEL_FAIL = 10002

	ERROR_ADD_CHANNEL_FAIL        = 10003
	ERROR_DELETE_CHANNEL_FAIL     = 10004
	ERROR_EDIT_CHANNEL_FAIL       = 10005
	ERROR_GET_CHANNEL_FAIL        = 10006
	ERROR_GET_CHANNEL_BY_ID_FAIL  = 10007
	ERROR_EXIST_CHANNEL_NAME      = 10008
	ERROR_EXIST_CHANNEL_NAME_FAIL = 10009

	ERROR_GET_CHANNEL_MESSAGE_FAIL = 20001

	ERROR_HASHPASSWORD      = 30001
	ERROR_EXIST_USER        = 30002
	ERROR_EXIST_EMAIL       = 30003
	ERROR_CHECK_EXIST_USER  = 30004
	ERROR_CHECK_EXIST_EMAIL = 30005
	ERROR_ADD_USER          = 30006
	ERROR_GET_USER          = 30008
	ERROR_NOT_EXIST_USER    = 30009

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 300010

	ERROR_GEN_TOKEN           = 40001
	ERROR_VERIFY_RFRESH_TOKEN = 40002
	ERROR_VERIFY_ACCESS_TOKEN = 40003

	ERROR_CAPTCHA_KEY_EMPTY  = 50001
	ERROR_CAPTCHA_VERIFYDOTS = 50002
	ERROR_CAPTCHA_STORE      = 50003

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 50004
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 50005
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 50006

	FAIL_CAPTCHA = 99999
)
