package main

import (
	"video/server/server"
	"video/server/socket"
)

func main() {
	serverIntf := new(server.Server)
	serverIntf.Ip = "127.0.0.1"
	serverIntf.Port = "5624"
	serverIntf.Intf = &socket.Socket{}
	serverIntf.StartServer()
}
