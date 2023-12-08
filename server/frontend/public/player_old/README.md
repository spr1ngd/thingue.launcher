## ThingUE消息收发nodejs代码示例

```javascript
const SingletonPixelStreaming = require('./Thing.UE/PixelStreaming');

const callback = function (params) {
    // 发送收到回复
    SingletonPixelStreaming.BroadcastMessageToFront("ToWebMessage", {type: "收到回复", data: params});
}
SingletonPixelStreaming.addMessageMap('ToUEMessage', async (params) => {
    console.log('getMessageFromWeb:', params);
    return await (callback && callback(params));
});
```

## 浏览器消息收发iframe测试代码

```html
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>iframe测试</title>
</head>

<body>
<iframe src="/static/player.html" id="player" width="530px" height="300px"></iframe>
<div>
    <select id="addr_select">
        <option value="/static/player.html">player.html</option>
        <option value="/static/player_v4.html">player_v4.html</option>
        <option value="/static/player_v5.html">player_v5.html</option>
    </select>
    <button onclick="sendMsg()">测试发消息</button>
</div>
<script>
    window.onload = function () {
        const addrSelect = document.getElementById("addr_select");
        addrSelect.addEventListener("change", function () {
            document.getElementById("player").src = addrSelect.value
        });
    }
    // 接收消息
    window.onmessage = async function (e) {
        console.log("主页面收到消息", e.data)
    }

    // 发送消息
    function sendMsg() {
        data = {text: "测试文本"}
        console.log("主页面发送消息", data)
        document.getElementById("player").contentWindow
                .postMessage(data)
    }
</script>
</body>
</html>
```
