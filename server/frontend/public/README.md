
## ThingUE消息收发nodejs代码示例

```javascript
const SingletonPixelStreaming = require('./Thing.UE/PixelStreaming');
setInterval(function () {
	console.log("发送测试消息")
	SingletonPixelStreaming.BroadcastMessageToFront("ToWebMessage", { timestamp: new Date() });
}, 5000);
const callback = function (params) {
	console.log(params)
}
SingletonPixelStreaming.addMessageMap('ToUEMessage', async (params) => {
	console.log('getMessageFromWeb:', params);
	return await (callback && callback(params));
});
```


## 浏览器消息收发iframe代码示例

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>iframe测试</title>
</head>
<body>
<iframe src="player_v5.html" id="player"></iframe>
<button onclick="sendMsg()">测试发消息</button>
<script>
    // 接收消息
    window.onmessage = async function (e) {
        console.log("主页面收到消息",e.data)
    }

    // 发送消息
    function sendMsg() {
        document.getElementById("player").contentWindow.postMessage({text: "测试文本"})
    }
</script>
</body>
</html>
```
