package main

import "video/client/handle"

func main() {
	client := handle.Client{
		Ip:   "192.168.96.131",
		Port: "56234",
	}
	client.InstallConnection()
}
