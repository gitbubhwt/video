package intf

import (
	"net"
	"video/common"
)

type ServerInterface interface {
	StartServer(ip, port string)                  //开启服务
	ProcessingConnection(conn net.Conn)           //处理连接
	CheckPackage(conn net.Conn) (bool, []byte)    //校验数据包
	ProcessingMsg(msg *common.Msg, conn net.Conn) //处理消息
}
