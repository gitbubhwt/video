package socket

import (
	"encoding/json"
	"errors"
	"net"
	"video/common"
	"video/intf"
	log "video/logger"
)

type Socket struct {
	Conn net.Conn
	intf.CommonSocket
}

//处理连接
func (this *Socket) ProcessingConnection(conn net.Conn) {
	for {
		this.Conn = conn
		isSuccess, data := this.CheckPackage()
		if !isSuccess {
			break
		}
		this.ProcessingMsg(data)
	}
	this.DeleteSession(conn)
}

//校验包
func (this *Socket) CheckPackage() (bool, []byte) {
	return common.ReceiveMsgByTcp(this.Conn, common.LOG_HEAD_SERVER)
}

//处理消息
func (this *Socket) ProcessingMsg(data []byte) {
	msg := new(common.Msg)
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(common.LOG_HEAD_SERVER, "json.Unmarshal fail", "err:", err, "data:", string(data))
		return
	}
	log.Info(common.LOG_HEAD_SERVER, "receive:", *msg)
	switch msg.MsgType {
	case common.MessageType_MSG_TYPE_HEART:
		{ //心跳
			this.ProcessingHeart(msg)
		}

	case common.MessageType_MSG_TYPE_VEDIO:
		{ //视频
			this.ProcessingVideo(msg, this.Conn)
		}
	}
}

//发送消息
func (this *Socket) SendMsg(conn net.Conn, msg *common.Msg) error {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Error(common.LOG_HEAD_WS_SERVER, "send msg fail,err:", err)
		return err
	}
	if b := common.SendMsgByTcp(data, conn, common.LOG_HEAD_SERVER); !b {
		err = errors.New(common.LOG_HEAD_WS_SERVER + " " + "send msg fail")
	}
	return err
}
