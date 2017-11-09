package wsocket

import (
	"net"
	"video/common"
	global "video/server/common"
)

//心跳
func (this *WSocket) ProcessingHeart(msg *common.Msg, conn net.Conn) {
	global.UpdateSessionMap(msg.From.Id, conn, common.SERVER_TYPE_WSOCKET)
}

//删除会话
func (this *WSocket) DeleteSession(conn net.Conn) {
	global.DeleteSessionMap(conn, common.SERVER_TYPE_WSOCKET)
}
