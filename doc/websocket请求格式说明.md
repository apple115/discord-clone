## 流程
用户连接：用户通过 WebSocket 连接到服务器。

身份验证：用户发送 connect 消息进行身份验证。

加入频道：用户发送 join_channel 消息加入频道。

发送消息：用户发送 send_message 消息到频道。

接收消息：用户接收频道中的消息。

心跳机制：用户定期发送 heartbeat 消息以保持连接。

断开连接：如果用户长时间没有心跳，服务器会关闭连接。

## WebSocket 请求格式说明
所有请求消息应为 JSON 格式，包含以下字段：
- type: 消息类型，可以是 "connect", "join_channel", "send_message", "heartbeat"
 - data: 消息数据，具体格式取决于消息类型
 示例：
connect
```json
 {
   "type": "connect",
   "data": {
     "token": "user_jwt_token",
     "user_id": "12345"
   }
 }
```
heartbeat
```json
 {
 		"type":"heartbeat"
 }
```
join_channel
```json
{
  "type": "join_channel",
  "data": {
    "channel_id": "your_channel_id"
  }
}
```
send_message
```
{
  "type": "send_message",
  "data": {
    "content": "Hello, world!"
  }
}
```
