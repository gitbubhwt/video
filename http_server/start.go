package main

import (
	"net/http"
	"video/common"
	webCommon "video/http_server/common"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

func main() {
<<<<<<< HEAD
	addr := ":" + common.WEB_SERVER_PORT
	http.Handle(webCommon.WEN_SERVER_STATIC_FILE_PATTERN, http.FileServer(http.Dir(webCommon.WEN_SERVER_STATIC_FILE_PATH)))
=======
	addr := ":5020"
	//log.Info(http.Dir("/http_server/html/video.html"))
	http.Handle("/", http.FileServer(http.Dir("./http_server")))
>>>>>>> 0f4fcefe0d4d408eeba8cbd981cabc427afd5dd5
	http.HandleFunc(route.ROUTE_HEAD_REQUEST, handle_video.VideoHeadHtml)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("Start web server fail,err:", err)
	}
}
