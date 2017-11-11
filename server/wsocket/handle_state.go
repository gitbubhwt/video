package wsocket

import (
	"video/common"
	global "video/server/common"
	log"video/logger"
)

//视频状态
func (this *WSocket) ProcessingVideoState(msg *common.Msg) {
	log.Info(common.LOG_HEAD_WS_SERVER, "receive:", *msg)
	global.SingleSendMsg(msg, this, common.SERVER_TYPE_WSOCKET)
}
