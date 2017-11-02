package wsocket

import (
	"encoding/json"
	"net"
	"video/common"
	"video/intf"
	log "video/logger"
)

type WSocket struct {
	intf.CommonSocket
	WsSocket *WsSocket
}

//处理连接
func (this *WSocket) ProcessingConnection(conn net.Conn) {
	buf := make([]byte, 1024)
	if n, err := conn.Read(buf); err != nil {
		log.Error("Ws hand shake read data fail,err:", err)
	} else {
		var isHttp bool = false
		if string(buf[0:3]) == common.WS_HAND_SHAKE {
			isHttp = true
		}
		if !isHttp {
			return
		}
		wsocket := NewWsSocket(conn)
		isHand := wsocket.HandShake(buf[:n]) //先握手
		if isHand {
			this.WsSocket = wsocket
			for {
				isSuccess, data := this.CheckPackage()
				if !isSuccess {
					break
				}
				this.ProcessingMsg(data)
			}
		}
	}
}

//校验包
func (this *WSocket) CheckPackage() (bool, []byte) {
	data, err := this.WsSocket.Read()
	if err != nil {
		log.Info(err)
		return false, nil
	}
	return true, data
}

//处理消息
func (this *WSocket) ProcessingMsg(data []byte) {
	msg := new(common.Msg)
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(common.LOG_HEAD_WS_SERVER, "json.Unmarshal fail", "err:", err, "data:", string(data))
		return
	}
	log.Info(common.LOG_HEAD_WS_SERVER, "receive:", *msg, string(data), msg.MsgType)
	switch msg.MsgType {
	case common.MessageType_MSG_TYPE_HEART:
		{
			//心跳
			this.ProcessingHeart(msg, this.WsSocket.Conn)
		}
	case common.MessageType_MSG_TYPE_VEDIO:
		{
			//视频
		}
	}
}
