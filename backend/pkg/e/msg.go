package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
	INVALID_PARAMS:                 "请求参数错误",
	ERROR_NOT_EXIST_CHANNEL:        "频道不存在",
	ERROR_CHECK_EXIST_CHANNEL_FAIL: "确认频道存在失败",
	ERROR_ADD_CHANNEL_FAIL:         "频道添加失败",
	ERROR_DELETE_CHANNEL_FAIL:      "频道删除失败",
	ERROR_EDIT_CHANNEL_FAIL:        "频道编辑失败",
	ERROR_GET_CHANNEL_FAIL:         "获取频道失败",
	ERROR_GET_CHANNEL_BY_ID_FAIL:   "通过ID获取频道失败",
	ERROR_GET_CHANNEL_MESSAGE_FAIL: "获取频道历史消息失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
