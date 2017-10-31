package main

import (
	"video/server/socket"
)

func main() {




	server := new(socket.Server)
	server.StartServer("127.0.0.1", "5664")
}
