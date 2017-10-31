package socket

import (
	"encoding/json"
	"net"
	"video/common"
	"video/intf"
	log "video/logger"
)

type Socket struct {
	intf.CommonSocket
}

//处理连接
func (this *Socket) ProcessingConnection(conn net.Conn) {
	for {
		isSuccess, data := this.CheckPackage(conn)
		if !isSuccess {
			break
		}
		this.ProcessingMsg(data, conn)
	}
	this.DeleteSession(conn)
}

//校验包
func (this *Socket) CheckPackage(conn net.Conn) (bool, []byte) {
	return common.ReceiveMsgByTcp(conn, common.LOG_HEAD_SERVER)
}

//处理消息
func (this *Socket) ProcessingMsg(data []byte, conn net.Conn) {
	msg := new(common.Msg)
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(common.LOG_HEAD_SERVER, "json.Unmarshal fail", "err:", err, "data:", string(data))
		return
	}
	log.Info(common.LOG_HEAD_SERVER,"receive:",*msg)
	switch msg.MsgType {
	case common.MessageType_MSG_TYPE_HEART:
		{ //心跳
			this.ProcessingHeart(msg, conn)
		}

	case common.MessageType_MSG_TYPE_VEDIO:
		{ //视频
			this.ProcessingVideo(msg, conn)
		}
	}
}
