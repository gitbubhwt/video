package main

import (
	"video/server/server"
	"video/server/socket"
)

func main() {
	serverIntf := new(server.Server)
	serverIntf.Ip = "127.0.0.1"
	serverIntf.Port = "56234"
	//mySocket:=new(socket.Socket)
	serverIntf.Intf = &socket.Socket{}
	serverIntf.StartServer()
}
