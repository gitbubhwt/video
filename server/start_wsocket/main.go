package main

import (
	"video/server/server"
	"video/server/wsocket"
)

func main() {
	serverIntf := new(server.Server)
	serverIntf.Ip = "127.0.0.1"
	serverIntf.Port = "56234"
	serverIntf.Intf=&wsocket.WSocket{}
	serverIntf.StartServer()
}
