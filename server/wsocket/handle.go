package wsocket

import (
	//"encoding/json"
	"net"
	//"video/common"
	//log "video/logger"
	"video/intf"
)

type WSocket struct {
	intf.CommonSocket
}

//处理连接
func (this *WSocket) ProcessingConnection(conn net.Conn) {
	//for {
	//	isSuccess, data := this.CheckPackage(conn)
	//	if !isSuccess {
	//		break
	//	}
	//	msg := new(common.Msg)
	//	err := json.Unmarshal(data, &msg)
	//	if err != nil {
	//		log.Error("json.Unmarshal data failed", err, string(data))
	//		break
	//	}
	//	this.ProcessingMsg(msg, conn)
	//}
	//DeleteSession(conn)
}
