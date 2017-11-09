var address ="192.168.96.131:56234";
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
        var obj=new Object();
        obj.name="ddddddddd";
        sendMsg(MessageType_MSG_TYPE_VEDIO,obj);
    }, 1000*60);
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
    ws.send(JSON.stringify(obj));
}
