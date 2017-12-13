package handle_video

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
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
	index := r.FormValue(webCommon.HEAD_VIDEO_ID)    //video id
	order := r.FormValue(webCommon.HEAD_VIDEO_ORDER) //video order
	log.Info("input params:", webCommon.HEAD_VIDEO_ID, ":", index, webCommon.HEAD_VIDEO_ORDER, ":", order)
	//视频封面
	video := new(common.MonVideo)
	mongo := db.GetMongo()
	orderT, _ := strconv.Atoi(order)
	mongo.C(common.MONGO_COLLECTION_VIDEO).Find(bson.M{"id": index, "path.orderNum": orderT}).One(&video)
	webCommon.GoToPage(w, route.ROUTE_PLAY_HTML_PATH, video)
}

//视频首页
func VideoIndexHtml(w http.ResponseWriter, r *http.Request) {
	videos := make([]common.MonVideo, 0)
	mongo := db.GetMongo()
	filter := bson.M{"path": 0}
	mongo.C(common.MONGO_COLLECTION_VIDEO).Find(nil).Select(filter).All(&videos)
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
	videos := make([]common.MonVideo, 0)
	mongo := db.GetMongo()
	query := mongo.C(common.MONGO_COLLECTION_VIDEO).Find(nil)
	pageOption := new(webCommon.MongoPageOption)
	pageOption.PageNo = 1
	pageOption.PageSize = 10
	query.Skip((pageOption.PageNo - 1) * pageOption.PageSize).Limit(pageOption.PageSize).All(&videos)
	//分页
	if page, err := pageOption.GetMongoPageOption(query, videos); err != nil {
		log.Error("Get page option fail,err:", err)
	} else {
		pageOption.List = videos
		webCommon.SendResponse(w, page)
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
	var filePath string
	rootPathT := db.GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_WEB_SERVER_PATH)
	if rootPath, ok := rootPathT.(string); ok {
		filePath = fmt.Sprintf(rootPath+webCommon.WEB_SERVER_UPLOAD_FILE_PATH, path)
	}
	if err := common.CreateFile(filePath, fileName, uploadFile); err != nil {
		msg = fmt.Sprintf("Video upload file fail,err:%v", err)
		log.Error(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
		return
	}
	msg = fmt.Sprintf("/upload/%s", path+fileName)
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
	video := new(common.MonVideo)
	video.Id = common.UniqueId()
	video.Name = videoName   //名称
	video.Cover = videoCover //封面
	video.Type = videoType   //类型
	video.CreateTime = time.Now().UnixNano() / 1e6
	size := 1 + len(videoChildFile)
	videoPaths := make([]common.MonVideoPath, size)
	for i := 0; i < size; i++ {
		videoPath := new(common.MonVideoPath)
		videoPath.OrderNum = i + 1
		if i == 0 {
			videoPath.Path = videoFile
		} else {
			videoPath.Path = videoChildFile[i-1]
		}
		videoPaths[i] = *videoPath
	}
	video.Path = videoPaths
	mongo := db.GetMongo()
	collection := mongo.C(common.MONGO_COLLECTION_VIDEO)
	err := collection.Insert(&video)
	var msg string
	if err != nil {
		msg := fmt.Sprintf("Video save fail,err:%v", err)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg)
	} else {
		msg := fmt.Sprintf("Video save success")
		webCommon.GoToResponse(w, common.ACK_SUCCESS, msg)
	}
	log.Info(msg)
}
