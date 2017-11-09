package handle

import (
	"encoding/json"
	"net"
	"time"
	"video/common"
	log "video/logger"
)

type Client struct {
	Ip   string
	Port string
}

var closeChan chan bool

//建立连接
func (this *Client) InstallConnection() {
	network := this.Ip + ":" + this.Port
	conn, err := net.Dial(common.SERVER_NET, network)
	if err != nil {
		log.Info("Connect server", network, "fail", err)
		return
	}
	closeChan = make(chan bool, 1)
	this.SendHeartPackage(conn)
	go this.ReceiveDataPackage(conn)
	<-closeChan
}

//接收数据包
func (this *Client) ReceiveDataPackage(conn net.Conn) {
	for {
		isSuccess, data := common.ReceiveMsgByTcp(conn, common.LOG_HEAD_CLIENT)
		if !isSuccess {
			break
		}
		log.Info(common.LOG_HEAD_CLIENT, "receive data:", data)
	}
	closeChan <- true
}

//处理消息
func (this *Client) ProcessingMsg(data []byte, conn net.Conn) {
	msg := new(common.Msg)
	err := json.Unmarshal(data, &msg)
	if err != nil {
		log.Error(common.LOG_HEAD_CLIENT, "json.Unmarshal fail", "err:", err, "data:", string(data))
		return
	}
	switch msg.MsgType {
	case common.MessageType_MSG_TYPE_VEDIO:
		{
			//视频
			this.ProcessingVideo(msg, conn)
		}
	}
}

//发送心跳包
func (this *Client) SendHeartPackage(conn net.Conn) {
	msg := common.Msg{
		MsgCommon: common.MsgCommon{
			Id:         "1",
			Ip:         "192.168.96.131",
			CreateTime: time.Now().Unix(),
		},
		MsgType: common.MessageType_MSG_TYPE_HEART,
	}
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Error(common.LOG_HEAD_CLIENT, "json.Marshal fail", "err:", err)
		return
	}
	var sendSuccess bool
	go func() {
		for {
			sendSuccess = common.SendMsgByTcp(bytes, conn, common.LOG_HEAD_CLIENT)
			if !sendSuccess {
				break
			}
			time.Sleep(common.HEART_RATE * time.Second)
		}
		closeChan <- true
	}()
}

//发送数据包
func (this *Client) SendDataPackage(msg *common.Msg, conn net.Conn) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Error(common.LOG_HEAD_CLIENT, "json.Marshal fail", "err:", err)
		return
	}
	common.SendMsgByTcp(bytes, conn, common.LOG_HEAD_CLIENT)
}
