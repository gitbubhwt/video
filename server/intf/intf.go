package intf

import (
	"net"
	"video/common"
)

type ServerInterface interface {
	ProcessingConnection(conn net.Conn)             //处理连接
	CheckPackage(conn net.Conn) (bool, []byte)      //校验数据包
	ProcessingMsg(msg *common.Msg, conn net.Conn)   //处理消息
	ProcessingHeart(msg *common.Msg, conn net.Conn) //处理心跳
	DeleteSession(conn net.Conn)                    //删除会话
}

type CommonSocket struct {
	Ip   string
	Port string
}
