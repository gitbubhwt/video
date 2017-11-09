var video = document.getElementById("Video1");
if(video.canPlayType){
    //暂停事件
    video.addEventListener("pause", function () {
        videoPauseAction();
    }, false);
    //播放事件
    video.addEventListener("playing", function () {
        videoPlayAction();
    }, false);

}else{
    alert("您的浏览器不支持播放");
}
//播放事件
function videoPlayAction(){
    logInfo("video play");
}
//暂停事件
function videoPauseAction(){
    logInfo("video pause");
}

