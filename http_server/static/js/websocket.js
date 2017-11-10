var address ="127.0.0.1:5624";
try {
    if ("WebSocket" in window) {
        ws = new WebSocket("ws://" + address);
    } else if ("MozWebSocket" in window) {
        ws = new MozWebSocket("ws://" + address);
    }
} catch (ex) {
    alert("连接异常");
}
//ws 开启
ws.onopen=function () {
    setInterval(function() {
        logInfo(ws,"open");
        var from=new Object();
        from.id="1";
        var to=new Object();
        to.id="1";
        var content="heart";
        sendMessage(from,to,MessageType_MSG_TYPE_HEART,content);
    }, 2000);
}
//关闭
ws.onclose=function(ev) {
    logInfo(ws,"close",ev);
}
//异常
ws.onerror=function (ev) {
    logInfo(ws,"close",ev);
}
//接收消息
ws.onmessage=function (ev) {
    logInfo(ws,"receive",ev);
}
//发送消息
function sendMsg(msgType,data) {
    var obj=new Object();
    obj.id="1";
    obj.ip="192.168.96.131";
    obj.msgType=msgType;
    obj.msgData=data;
    var sendData=JSON.stringify(obj);
    logInfo("ws send:",sendData);
    ws.send(sendData);
}

//发送消息
function sendMessage(from,to,msgType,content) {
    var data=new Object();
    data.from=from;
    data.to=to;
    data.content=content;
    data.msgType=msgType;
    var sendData=JSON.stringify(data);
    logInfo("ws send:",sendData);
    ws.send(sendData);
}