var address ="127.0.0.1:5624";
var websocket=null;

init();

//发送消息
function sendMessageFrom(from,to,msgType,content) {
    var data=new Object();
    data.from=from;
    data.to=to;
    data.content=content;
    data.msgType=msgType;
    return doSend(data);
}

//发送消息
function sendMessage(to,msgType,content) {
    var data=new Object();
    var from=new Object();
    from.id=USER_FROM_ID;
    data.from=from;
    data.to=to;
    data.content=content;
    data.msgType=msgType;
    return doSend(data);
}

//发送消息
function doSend(data) {
    var message=JSON.stringify(data);
    var success=false;
    websocket.send(message);
    for(var i=1;i<=3;i++){//失败后,尝试3次发送
        if(websocket.readyState==1){
            if(data.msgType!=MessageType_MSG_TYPE_HEART){
                logInfo("Ws send message success,data:",message);
            }
            success=true;
            break
        }else{
            websocket.send(message);
            logInfo("Ws try again send message,num:",i);
        }
    }
    return success;
}

function init() {
    try {
        if ("WebSocket" in window) {
            websocket = new WebSocket("ws://" + address);
        } else if ("MozWebSocket" in window) {
            websocket = new MozWebSocket("ws://" + address);
        }
    } catch (ex) {
        alert("连接异常");
    }
}
//连接打开
websocket.onopen = function (ev) {
    logInfo(websocket,"open",ev);
    var to=new Object();
    to.id="1";
    var content="";
    var success=sendMessage(to,MessageType_MSG_TYPE_HEART,content);
    var timer=setInterval(function() {
        success=sendMessage(to,MessageType_MSG_TYPE_HEART,content);
        if(!success){
            clearInterval(timer);//清除定时器
        }
    }, 5000);
};
//关闭
websocket.onclose=function(ev) {
    logInfo("Ws close",ev);
}
//异常
websocket.onerror=function (ev) {
    logInfo("Ws error",ev);
}
//接收消息
websocket.onmessage=function (ev) {
    logInfo("Ws receive data:",ev.data);
}

