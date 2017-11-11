var MessageType_MSG_TYPE_HEART = 1 ; //心跳
var MessageType_MSG_TYPE_VEDIO = 2 ;//视频
var MessageType_MSG_TYPE_VEDIO_STATE =3;//视频状态


var VIDEO_PAUSE ="pause" ;//暂停
var VIDEO_PLAY ="playing" ;//正在播放
var VIDEO_ENDED="ended"; //播放完毕
var VIDEO_TIME_UPDATE="timeupdate" ;//播放时间更新

var VIDEO_TIME_SEND_TIME=60*10 ;  //发送时间间隔

var USER_FROM_ID ="1" //用户id
var VIDEO_PLAY_TO_ID ="1" //视频播放接受用户id
var VIDEO_PAUSE_TO_ID ="1" //视频暂停接受用户id
var VIDEO_UPDATE_TIME_TO_ID ="" //视频更新时间接受用户id
var VIDEO_ENDED_TO_ID ="1" //视频结束接受用户id
// var VIDEO_PLAY_TO_ID ="1" //视频播放接受用户id