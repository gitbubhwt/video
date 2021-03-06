package socket

import (
	"net"
	"video/common"
	global "video/server/common"
)

//心跳
func ProcessingHeart(msg *common.Msg, conn net.Conn) {
	global.UpdateSessionMap(msg.Id, conn, common.SERVER_TYPE_SOCKET)
}

//删除会话
func DeleteSession(conn net.Conn) {
	global.DeleteSessionMap(conn, common.SERVER_TYPE_WSOCKET)
}
