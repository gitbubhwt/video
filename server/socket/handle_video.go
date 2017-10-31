package socket

import (
	"encoding/json"
	"net"
	"video/common"
	log "video/logger"
)

//处理视频
func (this *Socket)ProcessingVideo(msg *common.Msg, conn net.Conn) {
	log.Info(msg)
	msgData := msg.MsgData
	video := new(common.Video)
	err := json.Unmarshal(msgData, &video)
	if err != nil {
		log.Error("json.Unmarshal video failed", err)
		return
	}
	log.Info(video)
}
