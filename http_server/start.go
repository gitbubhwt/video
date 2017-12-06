package main

import (
	"net/http"
	"os"
	"strings"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

var webServerPath string

func main() {
	addr := ":" + common.WEB_SERVER_PORT

	staticPath := webServerPath + webCommon.WEN_SERVER_STATIC_PATH
	http.Handle(webCommon.WEB_SERVER_CSS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_JS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_IMG, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_UPLOAD, http.FileServer(http.Dir(staticPath)))
	http.HandleFunc(route.ROUTE_PLAY_REQUEST, handle_video.VideoPlayHtml)
	http.HandleFunc(route.ROUTE_INDEX_REQUEST, handle_video.VideoIndexHtml)
	http.HandleFunc(route.ROUTE_VIDEO_ADD_REQUEST, handle_video.VideoAddHtml)    //视频添加
	http.HandleFunc(route.ROUTE_VIDEO_LIST_REQUEST, handle_video.VideoListHtml)  //视频列表页面
	http.HandleFunc(route.ROUTE_VIDEO_UPLOAD_REQUEST, handle_video.VideoUpload)  //视频上传
	http.HandleFunc(route.ROUTE_VIDEO_SAVE_REQUEST, handle_video.VideoSave)      //视频保存
	http.HandleFunc(route.ROUTE_VIDEO_LIST_DATA_REQUEST, handle_video.VideoList) //视频列表数据
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
	webServerPath = common.GetCurrentDirectory()
	if strings.Index(webServerPath, "/") != -1 {
		arr := strings.Split(webServerPath, "/")
		path := ""
		for i := 0; i < len(arr)-1; i++ {
			path += arr[i] + "/"
		}
		mp := make(map[string]interface{})
		mp[common.SYSTEM_CONFIG_ROOT_PATH] = path
		mp[common.SYSTEM_CONFIG_WEB_SERVER_PATH] = webServerPath
		isSuccess := db.UpdateBatchHash(common.SYSTEM_CONFIG_KEY, mp)
		if !isSuccess {
			os.Exit(1)
			return
		}
	}
	err = db.GetMysqlDb()
	if err != nil {
		log.Error("fail to connect to mysql,err:", err)
		os.Exit(1)
		return
	}
}
