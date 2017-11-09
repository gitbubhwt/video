var video = document.getElementById("video");
var source = document.getElementById("source");
var name=source.src;
name="demo.mp4";
var vTime =0;
if(video.canPlayType){
    //暂停事件
    video.addEventListener(VIDEO_PAUSE, function () {
        videoPauseAction();
    }, false);
    //播放事件
    video.addEventListener(VIDEO_PLAY, function () {
        videoPlayAction();
    }, false);
    //播放完毕
    video.addEventListener(VIDEO_ENDED, function () {
        videoEnded();
    }, false);
    //时间更新
    var i=0;
    video.addEventListener(VIDEO_TIME_UPDATE, function () {
        vTime = video.currentTime;
        if(i==VIDEO_TIME_SEND_TIME){
            videoTimeUpdate(vTime);
            i=0;
        }
        i++;
    }, false);
}else{
    alert("您的浏览器不支持播放");
}
//播放事件
function videoPlayAction(){
    logInfo("video play");
    var obj=new Object();
    obj.name=name;
    obj.state=VIDEO_PLAY;
    obj.currentTime=vTime;
    sendMsg(MessageType_MSG_TYPE_VEDIO,obj);
}
//暂停事件
function videoPauseAction(){
    logInfo("video pause");
    var obj=new Object();
    obj.name=name;
    obj.state=VIDEO_PAUSE;
    obj.currentTime=vTime;
    sendMsg(MessageType_MSG_TYPE_VEDIO,obj);
}
//播放完毕
function videoEnded(){
    logInfo("video ended");
    var obj=new Object();
    obj.name=name;
    obj.state=VIDEO_ENDED;
    obj.currentTime=vTime;
    sendMsg(MessageType_MSG_TYPE_VEDIO,obj);
}
//时间更新
function videoTimeUpdate(vTime){
    logInfo("video time update is:",vTime);
}