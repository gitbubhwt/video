package common

import (
	"fmt"
	"net"
	log "video/logger"
)

//通过Tcp发送消息
func SendMsgByTcp(bytes []byte, conn net.Conn, logHead string) bool {
	data := make([]byte, 0)
	//封装包头
	head := []byte{PACKAGE_HEAD_A, PACKAGE_HEAD_B}
	data = append(data, head...)
	//封装数据长度
	util := new(Util)
	length := util.Uint642Bytes(uint64(len(bytes)))
	data = append(data, length...)
	//封装数据
	data = append(data, bytes...)
	//封装包尾
	data = append(data, PACKAGE_END)
	if _, err := conn.Write(data); err != nil {
		log.Error(logHead, "Send msg by tcp failed", "conn:", conn, "err", err)
		return false
	}
	return true
}

//通过Tcp接收消息
func ReceiveMsgByTcp(conn net.Conn, logHead string) (bool, []byte) {
	//校验包头
	head := make([]byte, 2)
	n, err := conn.Read(head)
	if err != nil || len(head) != n {
		log.Error(logHead, "read head fail", "err:", err, "size:", len(head), "read size:", n)
		return false, nil
	}
	if head[0] != PACKAGE_HEAD_A || head[1] != PACKAGE_HEAD_B {
		log.Error(logHead, "head data wrong", fmt.Sprintf("head:%x %x", head[0], head[1]))
		return false, nil
	}
	//获取包长度
	lenPackage := make([]byte, 8)
	n, err = conn.Read(lenPackage)
	if err != nil || len(lenPackage) != n {
		log.Error(logHead, "read length fail", "err:", err, "size:", len(lenPackage), "read size:", n)
		return false, nil
	}
	util := new(Util)
	length := util.Bytes2Uint64(lenPackage)
	//读取包数据
	data := make([]byte, length)
	n, err = conn.Read(data)
	if err != nil || len(data) != n {
		log.Error(logHead, "read data fail", "err:", err, "size:", len(data), "read size:", n)
		return false, nil
	}
	//校验包尾
	end := make([]byte, 1)
	n, err = conn.Read(end)
	if err != nil || len(end) != n {
		log.Error(logHead, "read end fail", "err:", err, "size:", len(data), "read size:", n)
		return false, nil
	}
	if end[0] != PACKAGE_END {
		log.Error(logHead, "head end wrong", fmt.Sprintf("end:%x", end[0]))
		return false, nil
	}
	return true, data
}
