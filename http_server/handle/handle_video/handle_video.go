package handle_video

import (
	"fmt"
	"net/http"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/route"
	log "video/logger"
)

//视频播放页面
func VideoPlayHtml(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := r.FormValue(webCommon.HEAD_VIDEO_ID)    //video id
	order := r.FormValue(webCommon.HEAD_VIDEO_ORDER) //video order
	log.Info("input params:", webCommon.HEAD_VIDEO_ID, ":", index, webCommon.HEAD_VIDEO_ORDER, ":", order)
	sqlDb := db.GetMysql()
	//视频封面
	video := new(common.Video)
	sqlDb.Id(index).Cols("cover").Get(video)
	//视频路径信息
	videoPath := new(common.VideoPath)
	sql := fmt.Sprintf(common.VIDEO_PAGE_SQL, index, order)
	sqlDb.Sql(sql).Get(videoPath)
	//播放视频信息
	videoPlay := new(webCommon.VideoPlay)
	videoPlay.VideoId = index
	videoPlay.Cover = video.Cover
	videoPlay.Path = videoPath.Path
	videoPlay.Order = order
	log.Info(*videoPlay)
	webCommon.GoToPage(w, route.ROUTE_PLAY_HTML_PATH, videoPlay)
}

//视频首页
func VideoIndexHtml(w http.ResponseWriter, r *http.Request) {
	videos := make([]common.Video, 0)
	videoPageSql := fmt.Sprintf(common.VIDEO_PAGE_LIST_SQL, common.DEFAULT_WHERE_SQL, "0", common.DEFAULT_PAGE_SIZE)
	sqlDb := db.GetMysql()
	sqlDb.Sql(videoPageSql).Find(&videos)
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

//视频数据列表
func VideoList(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	log.Info("Video list,input params is pageNo:", pageNo)
	videos := make([]common.Video, 0)
	sqlDb := db.GetMysql()
	sql := fmt.Sprintf(common.VIDEO_PAGE_LIST_SQL, common.DEFAULT_WHERE_SQL, pageNo, common.DEFAULT_PAGE_SIZE)
	sqlDb.Sql(sql).Find(&videos)
	//分页
	if pageOption, err := webCommon.GetPageOption(pageNo, common.DEFAULT_PAGE_SIZE, sql); err != nil {
		log.Error("Get page option fail,err:", err)
	} else {
		pageOption.List = videos
		webCommon.SendResponse(w, pageOption)
	}
}

//上传文件
func VideoUpload(w http.ResponseWriter, r *http.Request) {
	var msg string
	fileName := r.FormValue("name")
	path := r.FormValue("path")
	if fileName == common.STRING_NULL || path == common.STRING_NULL {
		msg = fmt.Sprintf("Video upload file fail,fileName,path is empty,fileName:%v,path:%v", fileName, path)
		log.Error(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
		return
	}
	//解析文件时候需要 ParseForm
	r.ParseForm()
	uploadFile, _, err := r.FormFile("file")
	if err != nil {
		msg = fmt.Sprintf("Video upload file fail,err:%v", err)
		log.Error(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
		return
	}
	rootPathT := db.GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_WEB_SERVER_PATH)
	if rootPath, ok := rootPathT.(string); ok {
		fileName = path + fileName
		path = fmt.Sprintf(rootPath+webCommon.WEB_SERVER_UPLOAD_FILE_PATH, fileName)
	}
	if err := common.CreateFile(path, uploadFile); err != nil {
		msg = fmt.Sprintf("Video upload file fail,err:%v", err)
		log.Error(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
		return
	}
	msg = fmt.Sprintf("/upload/%s", fileName)
	log.Info(msg)
	webCommon.GoToResponse(w, common.ACK_SUCCESS, msg)
}

//删除文件
func VideoDel(w http.ResponseWriter, r *http.Request) {
	var msg string
	path := r.FormValue("path")
	rootPathT := db.GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_WEB_SERVER_PATH)
	if rootPath, ok := rootPathT.(string); ok {
		path = rootPath + webCommon.WEN_SERVER_STATIC_PATH + path
	}
	if err := common.DelFile(path); err != nil {
		msg = fmt.Sprintf("Video del file fail,err:%v,path:%v", err, path)
		log.Error(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
	} else {
		msg = fmt.Sprintf("Video del file success,path:%v", path)
		log.Info(msg)
		webCommon.GoToResponse(w, common.ACK_SUCCESS, msg)
	}
}

//保存数据
func VideoSave(w http.ResponseWriter, r *http.Request) {
	videoName := r.FormValue("video_name")
	videoType := r.FormValue("video_type")
	videoCover := r.FormValue("video_cover")
	videoFile := r.FormValue("video_file")
	videoChildFile := r.PostForm["video_child_file"]
	log.Info(videoName, videoType, videoCover, videoFile, videoChildFile)
	sqlDb := db.GetMysql()
	video := new(common.Video)
	video.Name = videoName   //名称
	video.Cover = videoCover //封面
	video.Type = videoType   //类型

	//mongo := db.GetMongoDB()
	//collection := mongo.C(common.MONGO_COLLECTION_VIDEO)
	//collection.Insert(&video)
	var responseError error
	if _, err := sqlDb.InsertOne(video); err == nil {
		size := 1 + len(videoChildFile)
		videoPaths := make([]*common.VideoPath, size)
		for i := 0; i < size; i++ {
			videoPath := new(common.VideoPath)
			videoPath.VideoId = video.Id
			videoPath.OrderNum = 1
			if i == 0 {
				videoPath.Path = videoFile
			} else {
				videoPath.Path = videoChildFile[i-1]
			}
			videoPaths[i] = videoPath
		}
		if err := webCommon.BatchSaveVideoPath(videoPaths); err != nil {
			responseError = err
		}
	} else {
		responseError = err
	}
	if responseError != nil {
		msg := fmt.Sprintf("Video save fail,err:%v", responseError)
		log.Info(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
	} else {
		msg := fmt.Sprintf("Video save success")
		log.Info(msg)
		webCommon.GoToResponse(w, common.ACK_SUCCESS, msg)
	}
}
