package socket

import (
	"encoding/json"
	"net"
	"video/common"
	log "video/logger"
)

//开启服务
func (this *Server) StartServer(ip, port string) {
	addr := ip + ":" + port
	listen, err := net.Listen(common.SERVER_NET, addr)
	if err != nil {
		log.Error("Start", addr, "server failed", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Server accept client failed", err)
			DeleteSession(conn) //删除会话
			break
		}
		go this.ProcessingConnection(conn)
	}
}

//处理连接
func (this *Server) ProcessingConnection(conn net.Conn) {
	for {
		isSuccess, data := this.CheckPackage(conn)
		if !isSuccess {
			break
		}
		msg := new(common.Msg)
		err := json.Unmarshal(data, &msg)
		if err != nil {
			log.Error("json.Unmarshal data failed", err, string(data))
			break
		}
		this.ProcessingMsg(msg, conn)
	}
	DeleteSession(conn)
}

//校验包
func (this *Server) CheckPackage(conn net.Conn) (bool, []byte) {
	//校验包头
	head := make([]byte, 2)
	n, err := conn.Read(head)
	if err != nil || len(head) != n {
		log.Error("read head package failed or head package length is wrong", err, len(head), n)
		return false, nil
	}
	if head[0] != common.PACKAGE_HEAD_A || head[1] != common.PACKAGE_HEAD_B {
		log.Error("head package is wrong", head[0], head[1])
		return false, nil
	}
	//获取包长度
	lenPackage := make([]byte, 8)
	n, err = conn.Read(lenPackage)
	if err != nil || len(lenPackage) != n {
		log.Error("read len package failed or len package length is wrong", err, len(lenPackage), n)
		return false, nil
	}
	util := new(common.Util)
	length := util.Bytes2Uint64(lenPackage)
	//读取包数据
	data := make([]byte, length)
	n, err = conn.Read(data)
	if err != nil || len(data) != n {
		log.Error("read data package failed or data package length is wrong", err, len(data), n)
		return false, nil
	}
	//校验包尾
	end := make([]byte, 1)
	n, err = conn.Read(end)
	if err != nil || len(end) != n {
		log.Error("read end package failed or end package length is wrong", err, len(end), n)
		return false, nil
	}
	if end[0] != common.PACKAGE_END {
		log.Error("end package is wrong", head[0], head[1])
		return false, nil
	}
	return true, data
}

//处理消息
func (this *Server) ProcessingMsg(msg *common.Msg, conn net.Conn) {
	switch msg.MsgType {
	case common.MessageType_MSG_TYPE_HEART:
		{ //心跳
			ProcessingHeart(msg, conn)
		}

	case common.MessageType_MSG_TYPE_VEDIO:
		{ //视频
			ProcessingVideo(msg, conn)
		}
	}
}
