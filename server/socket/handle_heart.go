package socket

import (
	"net"
	"video/common"
	global "video/server/common"
)

//心跳
func (this *Socket)ProcessingHeart(msg *common.Msg) {
	global.UpdateSessionMap(msg.From.Id, this.Conn, common.SERVER_TYPE_SOCKET)
}

//删除会话
func (this *Socket)DeleteSession(conn net.Conn) {
	global.DeleteSessionMap(conn, common.SERVER_TYPE_SOCKET)
}
