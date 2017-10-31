package server

import (
	"net"
	"video/common"
	log "video/logger"
	"video/server/intf"
)

type Server struct {
	Ip   string
	Port string
	Intf intf.ServerInterface
}

//开启服务
func (this *Server) StartServer() {
	addr := this.Ip + ":" + this.Port
	listen, err := net.Listen(common.SERVER_NET, addr)
	if err != nil {
		log.Error("Start", addr, "server failed", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("Server accept client failed", err)
			this.Intf.DeleteSession(conn) //删除会话
			break
		}
		go this.Intf.ProcessingConnection(conn)
	}
}
