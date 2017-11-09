package wsocket

import (
	"encoding/json"
	"net"
	"video/common"
	"video/intf"
	log "video/logger"
	"fmt"
)

type WSocket struct {
	intf.CommonSocket
	WsSocket *WsSocket
}

//处理连接
func (this *WSocket) ProcessingConnection(conn net.Conn) {
	wsocket := NewWsSocket(conn)
	this.WsSocket = wsocket
	var isHand bool
	for {
		opcodeByte := make([]byte, 1)
		_, err := this.WsSocket.Conn.Read(opcodeByte)
		if err != nil {
			break
		}
		this.WsSocket.OpcodeByte = opcodeByte
		if opcodeByte[0] == common.WS_ON_LINE {
			isHand = this.UserOnline()
		} else if opcodeByte[0] == common.WS_NORMAL {
			if !isHand {
				break
			}
			isSuccess, data := this.CheckPackage()
			if !isSuccess {
				continue
			}
			this.ProcessingMsg(data)
		} else if opcodeByte[0] == common.WS_OFF_LINE {
			this.UserOffline()
		}
	}
}

//用户结束
func (this *WSocket) UserOffline() bool {
	log.Info("User offline")
	data, err := this.WsSocket.Read()
	if err != nil {
		log.Error("User offline,read data fail", err)
		return false
	}
	str := string(data)
	if err = this.WsSocket.Write([]byte(str)); err != nil {
		log.Error("User offline,write data to user fail", err)
		return false
	}
	return true
}

//用户登录
func (this *WSocket) UserOnline() bool {
	log.Info("User online")
	isHand := this.WsSocket.HandShake() //先握手
	return isHand
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
