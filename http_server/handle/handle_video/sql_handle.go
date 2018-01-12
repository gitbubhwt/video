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
func VideoPlayHtml_1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	index := r.FormValue(webCommon.HEAD_VIDEO_ID)    //video id
	order := r.FormValue(webCommon.HEAD_VIDEO_ORDER) //video order
	log.Info("input params:", webCommon.HEAD_VIDEO_ID, ":", index, webCommon.HEAD_VIDEO_ORDER, ":", order)
	sqlDb := db.GetMysql()
	//视频封面
	video := new(common.MonVideo)
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
	webCommon.GoToPage(w, route.ROUTE_play_html, videoPlay)
}
//视频首页
func VideoIndexHtml_1(w http.ResponseWriter, r *http.Request) {
	videos := make([]common.Video, 0)
	videoPageSql := fmt.Sprintf(common.VIDEO_PAGE_LIST_SQL, common.DEFAULT_WHERE_SQL, "0", common.DEFAULT_PAGE_SIZE)
	sqlDb := db.GetMysql()
	sqlDb.Sql(videoPageSql).Find(&videos)
	webCommon.GoToPage(w, route.ROUTE_admin_html, videos)
}


//视频数据列表
func VideoList_1(w http.ResponseWriter, r *http.Request) {
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

//保存数据
func VideoSave_1(w http.ResponseWriter, r *http.Request) {
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
	video.CreateTime = time.Now()

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
		video.Path = videoPaths
		mongo := db.GetMongo()
		collection := mongo.C(common.MONGO_COLLECTION_VIDEO)
		err = collection.Insert(&video)
		if err := webCommon.BatchSaveVideoPath(videoPaths); err != nil {
			responseError = err
		}
	} else {
		responseError = err
	}
	if responseError != nil {
		msg := fmt.Sprintf("Video save fail,err:%v", responseError)
		log.Info(msg)
		webCommon.GoToResponse(w, common.ACK_FAIL, msg,nil)
	} else {
		msg := fmt.Sprintf("Video save success")
		log.Info(msg)
		webCommon.GoToResponse(w, common.ACK_SUCCESS, msg,nil)
	}
}