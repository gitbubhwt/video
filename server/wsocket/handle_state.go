package wsocket

import (
	"video/common"
	global "video/server/common"
)
//视频状态
func (this *WSocket) ProcessingVideoState(msg *common.Msg) {
	global.UpdateSessionMap(msg.From.Id, this.WsSocket.Conn, common.SERVER_TYPE_WSOCKET)
}
