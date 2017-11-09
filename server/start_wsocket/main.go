package main

import (
	"video/server/server"
	"video/server/wsocket"
	"log"
)

func main() {
	log.Println("ddddd")
	serverIntf := new(server.Server)
	serverIntf.Ip = "127.0.0.1"
	serverIntf.Port = "5624"
	serverIntf.Intf=&wsocket.WSocket{}
	serverIntf.StartServer()
}
