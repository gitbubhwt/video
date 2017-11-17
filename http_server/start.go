package main

import (
	"net/http"
	"os"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

var rootPath string

func main() {
	addr := ":" + common.WEB_SERVER_PORT

	staticPath := rootPath + webCommon.WEN_SERVER_STATIC_PATH
	http.Handle(webCommon.WEB_SERVER_CSS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_JS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_IMG, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_UPLOAD, http.FileServer(http.Dir(staticPath)))
	http.HandleFunc(route.ROUTE_PLAY_REQUEST, handle_video.VideoPlayHtml)
	http.HandleFunc(route.ROUTE_INDEX_REQUEST, handle_video.VideoIndexHtml)
	http.HandleFunc(route.ROUTE_VIDEO_ADD_REQUEST, handle_video.VideoAddHtml)   //视频添加
	http.HandleFunc(route.ROUTE_VIDEO_LIST_REQUEST, handle_video.VideoListHtml) //视频列表
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("Start web server fail,err:", err)
	}
}

func init() {
	err := db.GetRedisClient()
	if err != nil {
		log.Error("fail to connect to redis,err:", err)
		os.Exit(1)
		return
	}
	rootPath = common.GetCurrentDirectory()
	isSuccess := db.UpdateHash(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_ROOT_PATH, rootPath)
	if !isSuccess {
		os.Exit(1)
		return
	}
	_, err = db.GetOneDb()
	if err != nil {
		log.Error("fail to connect to mysql,err:", err)
		os.Exit(1)
		return
	}
}
