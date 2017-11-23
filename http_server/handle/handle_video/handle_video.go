package handle_video

import (
	"fmt"
	"net/http"
	"time"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/route"
	log "video/logger"
)

//视频播放页面
func VideoPlayHtml(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := r.Form.Get(webCommon.HEAD_VIDEO_INDEX) //video index
	log.Info("input params:", webCommon.HEAD_VIDEO_INDEX, ":", index)
	video := new(common.Video)
	video.Index = index
	video.ImgSrc = fmt.Sprintf("img/%s.jpg", index)
	video.Path = "upload/demo.mp4"
	webCommon.GoToPage(w, route.ROUTE_PLAY_HTML_PATH, video)
}

//视频首页
func VideoIndexHtml(w http.ResponseWriter, r *http.Request) {
	videos := make([]common.Video, 4)
	count := 1
	for i := 0; i < len(videos); i++ {
		video := new(common.Video)
		video.Index = fmt.Sprintf("%d", count)
		video.ImgSrc = fmt.Sprintf("img/%d.jpg", count)
		video.Name = fmt.Sprintf("电影%d", count)
		count++
		videos[i] = *video
	}
	webCommon.GoToPage(w, route.ROUTE_INDEX_HTML_PATH, videos)
}

//视频新增页面
func VideoAddHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_ADD_HTML_PATH, nil)
}

//视频列表页面
func VideoListHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_LIST_HTML_PATH, nil)
}

//上传文件
func VideoUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uploadFile, _, err := r.FormFile("file")
	if err != nil {
		log.Error("Video upload file fail,err:", err)
		return
	}
	rootPathT := db.GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_ROOT_PATH)
	if rootPath, ok := rootPathT.(string); ok {
		//path := fmt.Sprintf(rootPath+webCommon.WEB_SERVER_UPLOAD_FILE_PATH, time.Now().Unix())
		path := fmt.Sprintf(webCommon.WEB_SERVER_UPLOAD_FILE_TEMP_PATH, time.Now().Unix())
		if err := common.CreateFile(path, uploadFile); err != nil {
			log.Error("Video upload file fail,err:", err)
			return
		}
		log.Info("Video upload file success")
		w.Write([]byte("upload success"))
	} else {
		log.Error(common.SYSTEM_CONFIG_ROOT_PATH, "type is wrong", rootPath)
	}
}
