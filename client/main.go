package main

import "video/client/handle"

func main() {
	client := handle.Client{
		Ip:   "127.0.0.1",
		Port: "5624",
	}
	client.InstallConnection()
}
