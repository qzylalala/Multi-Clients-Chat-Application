let socket = new WebSocket("ws://localhost:8080/ws")

let connect = cb => {
    console.log("connecting");

    // 指定连接成功后的回调函数
    socket.onopen = () => {
        console.log("Successfully Connected");
    };

    // 指定收到服务器数据后的回调函数
    socket.onmessage = msg => {
        console.log(msg);
        cb(msg);
    };

    // 指定连接关闭后的回调函数
    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    // 指定报错时的回调函数
    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };
};

let sendMsg = (msg) => {
    console.log("sending msg: ", msg);
    socket.send(msg);
};

export { connect, sendMsg };