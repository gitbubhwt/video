package common

import (
	"net"
	"sync"
	"video/common"
	"video/intf"
	log "video/logger"
)

var lock sync.Mutex                       //同步锁
var SocketSessionMap map[string]net.Conn  //socket会话
var WsocketSessionMap map[string]net.Conn //wSocket会话

//更新会话map
func UpdateSessionMap(id string, conn net.Conn, serverType int) {
	lock.Lock()
	defer lock.Unlock()
	switch serverType {
	case common.SERVER_TYPE_SOCKET:
		{
			if SocketSessionMap == nil { //初始化session
				SocketSessionMap = make(map[string]net.Conn)
			}
			if _, ok := SocketSessionMap[id]; !ok {
				SocketSessionMap[id] = conn //写入session
			}
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			if WsocketSessionMap == nil { //初始化session
				WsocketSessionMap = make(map[string]net.Conn)
			}
			if _, ok := WsocketSessionMap[id]; !ok {
				WsocketSessionMap[id] = conn //写入session
			}
		}
	}

}

//删除会话map
func DeleteSessionMap(conn net.Conn, serverType int) {
	lock.Lock()
	defer lock.Unlock()
	switch serverType {
	case common.SERVER_TYPE_SOCKET:
		{
			for k, v := range SocketSessionMap {
				if v == conn {
					delete(SocketSessionMap, k)
					break
				}
			}
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			for k, v := range WsocketSessionMap {
				if v == conn {
					delete(WsocketSessionMap, k)
					break
				}
			}
		}
	}
	if err := conn.Close(); err != nil {
		log.Error("close client user failed", err)
	}
}

//单一发送消息
func SingleSendMsg(msg *common.Msg, intf intf.ServerInterface, serverType int) bool {
	ok := true
	switch serverType {
	case common.SERVER_TYPE_WSOCKET:
		{
			if conn, ok1 := WsocketSessionMap[msg.To.Id]; ok1 {
				if err := intf.SendMsg(conn, msg); err != nil {
					ok = false
					log.Error(common.LOG_HEAD_WS_SERVER, "send msg fail,err:", err)
				}
			}
		}
	case common.SERVER_TYPE_SOCKET:
		{
			if conn, ok1 := SocketSessionMap[msg.To.Id]; ok1 {
				if err := intf.SendMsg(conn, msg); err != nil {
					ok = false
					log.Error(common.LOG_HEAD_SERVER, "send msg fail,err:", err)
				}
			}
		}
	}
	return ok
}
