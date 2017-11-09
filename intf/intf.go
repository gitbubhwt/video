package intf

import (
	"net"
	"video/common"
)

type ServerInterface interface {
	ProcessingConnection(conn net.Conn)             //处理连接
	CheckPackage() (bool, []byte)                   //校验数据包
	ProcessingMsg(data []byte)                      //处理消息
	ProcessingHeart(msg *common.Msg, conn net.Conn) //处理心跳
	DeleteSession(conn net.Conn)                    //删除会话
}

type CommonSocket struct {
	Ip   string
	Port string
}

type ClientInterface interface {
	InstallConnection()                             //建立连接
	ReceiveDataPackage(conn net.Conn)               //接收数据包
	ProcessingMsg(data []byte, conn net.Conn)       //处理消息
	SendDataPackage(msg *common.Msg, conn net.Conn) //发送数据包s
	SendHeartPackage(conn net.Conn)                 //发送心跳包
}
