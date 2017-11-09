package main

import (
	"video/server/server"
	"video/server/socket"
)

func main() {
	serverIntf := new(server.Server)
	serverIntf.Ip = "192.168.96.131"
	serverIntf.Port = "56234"
	serverIntf.Intf = &socket.Socket{}
	serverIntf.StartServer()
}
