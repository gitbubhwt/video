package main

import (
	"net/http"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

func main() {


	
	addr := ":5020"
	//log.Info(http.Dir("/http_server/html/video.html"))
	http.Handle("/", http.FileServer(http.Dir("./http_server")))
	http.HandleFunc(route.ROUTE_HEAD_REQUEST, handle_video.VideoHeadHtml)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("Start web server fail,err:", err)
	}
}
