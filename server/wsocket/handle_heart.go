package wsocket

import (
	"net"
	"video/common"
	global "video/server/common"
)

//心跳
func (this *WSocket) ProcessingHeart(msg *common.Msg) {
	global.UpdateSessionMap(msg.From.Id, this.WsSocket.Conn, common.SERVER_TYPE_WSOCKET)
	global.SingleSendMsg(msg, this, common.SERVER_TYPE_WSOCKET)
}

//删除会话
func (this *WSocket) DeleteSession(conn net.Conn) {
	global.DeleteSessionMap(conn, common.SERVER_TYPE_WSOCKET)
}
