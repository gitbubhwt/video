package common

import (
	"encoding/json"
	"net"
	"sync"
	"video/common"
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
			updateSessionMap(SocketSessionMap, id, conn)
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			updateSessionMap(WsocketSessionMap, id, conn)
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
			deleteSession(SocketSessionMap, conn)
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			deleteSession(WsocketSessionMap, conn)
		}
	}
}

//单一发送消息
func SingleSendMsg(msg common.Msg, id string, serverType int) bool {
	var ok bool = false
	switch serverType {
	case common.SERVER_TYPE_SOCKET:
		{
			ok = singleSend(SocketSessionMap, msg, id)
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			ok = singleSend(WsocketSessionMap, msg, id)
		}
	}
	return ok
}

//广播发送消息
func BroadCastSendMsg(msg common.Msg, serverType int) {
	switch serverType {
	case common.SERVER_TYPE_SOCKET:
		{
			broadCastSend(SocketSessionMap, msg)
		}
	case common.SERVER_TYPE_WSOCKET:
		{
			broadCastSend(WsocketSessionMap, msg)
		}
	}
}

//更新会话
func updateSessionMap(session map[string]net.Conn, id string, conn net.Conn) {
	if session == nil { //初始化session
		session = make(map[string]net.Conn)
	}
	if _, ok := session[id]; !ok {
		session[id] = conn //写入session
	}
}

//删除会话
func deleteSession(session map[string]net.Conn, conn net.Conn) {
	for k, v := range session {
		if v == conn {
			delete(session, k)
			break
		}
	}
	if err := conn.Close(); err != nil {
		log.Error("close client user failed", err)
	}
}

//单一发送
func singleSend(session map[string]net.Conn, msg common.Msg, id string) bool {
	if val, ok := session[id]; ok {
		bytes, err := json.Marshal(msg)
		if err != nil {
			log.Error("broad cast send,json.Marshal failed", err)
			return false
		}
		return common.SendMsgByTcp(bytes, val, common.LOG_HEAD_SERVER)
	}
	return false
}

//广播发送
func broadCastSend(session map[string]net.Conn, msg common.Msg) {
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Error("broad cast send,json.Marshal failed", err)
		return
	}
	for _, v := range session {
		common.SendMsgByTcp(bytes, v, common.LOG_HEAD_SERVER)
	}
}
