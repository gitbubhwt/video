package main

import (
	"net/http"
	"strings"
	"video/common"
	webCommon "video/http_server/common"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

func main() {
	addr := ":" + common.WEB_SERVER_PORT
	staticPath := webCommon.WEN_SERVER_STATIC_FILE_PATH
	directory := common.GetCurrentDirectory()
	if strings.Index(directory, webCommon.SERVER_ROOT_PATH) != -1 {
		staticPath = directory + staticPath
	} else {
		staticPath = "./" + webCommon.SERVER_ROOT_PATH + staticPath
	}
	http.Handle(webCommon.WEN_SERVER_STATIC_FILE_PATTERN, http.FileServer(http.Dir(staticPath)))
	http.HandleFunc(route.ROUTE_PLAY_REQUEST, handle_video.VideoPlayHtml)
	http.HandleFunc(route.ROUTE_HEAD_REQUEST, handle_video.VideoHeadHtml)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("Start web server fail,err:", err)
	}
}
