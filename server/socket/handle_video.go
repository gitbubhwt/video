package socket

import (
	"encoding/json"
	"net"
	"video/common"
	log "video/logger"
)

//服务器处理视频
func (this *Socket) ProcessingVideo(msg *common.Msg, conn net.Conn) {
	msgData := msg.MsgData
	video := new(common.VideoServer)
	err := json.Unmarshal(msgData, &video)
	if err != nil {
		log.Error("Processing video failed", err)
		return
	}
	path := common.SERVER_FILE_ROOT_PATH + video.Name
	common.WriteFile(path, video.Data, video.Off)
}
