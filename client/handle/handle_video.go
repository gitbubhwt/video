package handle

import (
	"encoding/json"
	"net"
	"video/common"
	log "video/logger"
)

//客户端处理视频
func (this *Client) ProcessingVideo(msg *common.Msg, conn net.Conn) {
	msgData,_ := msg.Content.([]byte)
	video := new(common.VideoClient)
	err := json.Unmarshal(msgData, &video)
	if err != nil {
		log.Error("Processing video fail,err:", err)
		return
	}
	path := common.CLIENT_FILE_ROOT_PATH + video.Name
	common.ReadFile(path, conn)
}
