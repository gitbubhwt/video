var video = document.getElementById("video");
var source = document.getElementById("source");
var progress;
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
        document.getElementById("videoProgress").value=vTime;
        if(i==VIDEO_TIME_SEND_TIME){
            videoTimeUpdate(vTime);
            i=0;
        }
        i++;
    }, false);
    video.addEventListener("loadedmetadata", function () {
            var totalTime=video.duration.toFixed(1);
			document.getElementById("videoProgress").max=totalTime;
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
    var to=new Object();
    to.id=VIDEO_PLAY_TO_ID;
    sendMessage(to,MessageType_MSG_TYPE_VEDIO_STATE,obj);
    document.getElementById("videoProgress").display="";
}
//暂停事件
function videoPauseAction(){
    logInfo("video pause");
    var obj=new Object();
    obj.name=name;
    obj.state=VIDEO_PAUSE;
    obj.currentTime=vTime;
    var to=new Object();
    to.id=VIDEO_PAUSE_TO_ID;
    sendMessage(to,MessageType_MSG_TYPE_VEDIO_STATE,obj);
}
//播放完毕
function videoEnded(){
    logInfo("video ended");
    var obj=new Object();
    obj.name=name;
    obj.state=VIDEO_ENDED;
    obj.currentTime=vTime;
    var to=new Object();
    to.id=VIDEO_ENDED_TO_ID;
    sendMessage(to,MessageType_MSG_TYPE_VEDIO_STATE,obj);
}
//时间更新
function videoTimeUpdate(vTime){
    logInfo("video time update is:",vTime);
}

var play =0;
//页面播放视频
function videoPlay(elem){
	switch(play){
		case 0:{
			elem.src="img/aoz.png";
		    video.play();
		    play=1;
			break;
		}
		case 1:{
			elem.src="img/aox.png";
		    video.pause();
		    play=0;
			break;
		}
	}
}
//进度条拖放，视频快进快退
var elem = document.querySelector('input[type="range"]');
var rangeValue = function(){
  var newValue = elem.value;
  video.currentTime= newValue;
};
elem.addEventListener("input", rangeValue);

