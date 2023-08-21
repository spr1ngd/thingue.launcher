import mitt from "mitt"

const origin = window.location.origin.replace('http://', 'ws://').replace('https://', 'wss://')
const path = window.location.pathname.slice(0, location.pathname.lastIndexOf("/"))
const wsURL = `${origin}${path}/ws/admin`

let socket;
let reconnectInterval = 1000; // 初始重连间隔，单位毫秒
const maxReconnectInterval = 60000; // 最大重连间隔，单位毫秒
let reconnectTimer;
const emitter = mitt();

function connectWebSocket() {
    socket = new WebSocket(wsURL); // 连接到WebSocket服务器

    socket.addEventListener('open', () => {
        console.log('WebSocket connected');
        reconnectInterval = 1000; // 重置重连间隔为初始值
    });

    socket.addEventListener('close', event => {
        console.log('WebSocket closed', event);
        // 触发重连逻辑
        reconnect();
    });

    socket.addEventListener('message', event => {
        console.log('WebSocket message received:', event.data);
        var msg = JSON.parse(event.data);
		if (msg.type === 'update') {
            emitter.emit("update", "")
        }
    });

    socket.addEventListener('error', error => {
        console.error('WebSocket error:', error);
    });
}

function reconnect() {
    clearTimeout(reconnectTimer);
    reconnectTimer = setTimeout(() => {
        reconnectInterval = Math.min(reconnectInterval * 2, maxReconnectInterval); // 使用指数退避算法递增重连间隔
        console.log(`Attempting to reconnect in ${reconnectInterval}ms`);
        connectWebSocket();
    }, reconnectInterval);
}

export {
    connectWebSocket, emitter
}