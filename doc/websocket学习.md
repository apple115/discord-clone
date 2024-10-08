

WebSocket API 可以在用户的浏览器和服务器之间打开双向交互式通信会话。使用此 API，您可以向服务器发送消息并接收响应，而无需轮询服务器以获取回复。
websocket 是双全工
```javascript
const httpsWebSocket = new WebSocket('wss://websocket.example.org');
console.log(httpsWebSocket.url); // 'wss://websocket.example.org'
... // Do something with socket
httpsWebSocket.close();
```

通过 一个 client的池 
