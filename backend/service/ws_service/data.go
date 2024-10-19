package ws_service

var responceData = map[string]map[string]interface{}{
	"error": {
		"type": "error",
		"data": "error",
	},
	"connect_success": {
		"type": "connect_ack",
		"data": map[string]string{
			"status":  "success",
			"message": "Connected successfully",
		},
	},
}
