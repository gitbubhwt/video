package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/handle/handle_index"
	"video/http_server/handle/handle_login"
	"video/http_server/handle/handle_video"
	"video/http_server/route"
	log "video/logger"
)

var webServerPath string

type Action func(writer http.ResponseWriter, request *http.Request)

var urlMaps map[string]Action

func main() {
	addr := ":" + common.WEB_SERVER_PORT

	staticPath := webServerPath + webCommon.WEB_SERVER_STATIC_PATH
	http.Handle(webCommon.WEB_SERVER_CSS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_JS, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_IMG, http.FileServer(http.Dir(staticPath)))
	http.Handle(webCommon.WEB_SERVER_UPLOAD, http.FileServer(http.Dir(staticPath)))
	for k, _ := range urlMaps {
		http.HandleFunc(k, filter)
	}
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error("Start web server fail,err:", err)
	}
}

//注册路由
func register() {
	if urlMaps == nil {
		urlMaps = make(map[string]Action)
	}
	urlMaps[route.ROUTE_play] = handle_video.VideoPlayHtml                  //播放页面
	urlMaps[route.ROUTE_index] = handle_video.VideoIndexHtml                //首页
	urlMaps[route.ROUTE_admin_video_add] = handle_video.AdminVideoAddHtml   //视频添加
	urlMaps[route.ROUTE_admin_video_list] = handle_video.AdminVideoListHtml //视频列表页面
	urlMaps[route.ROUTE_admin_video_upload] = handle_video.AdminVideoUpload //视频上传
	urlMaps[route.ROUTE_admin_video_save] = handle_video.AdminVideoSave     //视频保存
	urlMaps[route.ROUTE_admin_video_pageList] = handle_video.AdminVideoList //视频列表数据
	urlMaps[route.ROUTE_admin_tologin] = handle_login.AdminToLoginHtml      //管理员登陆页面
	urlMaps[route.ROUTE_admin_login] = handle_login.AdminLogin              //管理员登陆
	urlMaps[route.ROUTE_admin_index] = handle_index.AdminIndexHtml          //后台管理首页
}

//url过滤器
func filter(response http.ResponseWriter, request *http.Request) {
	url := request.RequestURI
	log.Info(fmt.Sprintf("Request url:%v", url))
	if action, ok := urlMaps[url]; ok {
		var requestIp string
		if strings.Index(request.RemoteAddr, ":") != -1 {
			arr := strings.Split(request.RemoteAddr, ":")
			requestIp = arr[0]
		}
		//用户登录
		if strings.Index(url, route.ROUTE_admin) == -1 {
			action(response, request) //执行方法
			return
		}
		//后端管理员登录
		if url == route.ROUTE_admin_login || url == route.ROUTE_admin_tologin { //登录相关请求
			action(response, request)
			return
		}
		ipKey := fmt.Sprintf(common.IP_SESSION_HASH_KEY, common.DB_admin)
		sid := db.GetStringValue(ipKey, requestIp)
		tologin, ok := urlMaps[route.ROUTE_admin_tologin]
		if sid == common.STRING_NULL && ok {
			log.Info(sid, ok)
			tologin(response, request)
			return
		}
		//判断会话是否还存在
		key := fmt.Sprintf(common.SESSION_HASH_KEY, common.DB_admin, sid)
		redis := db.GetRedisClient()
		state := redis.TTL(key).Val()
		if state == -2*time.Second { //不存在,跳转至登录
			log.Info(fmt.Sprintf("Current session is not exist,key:%v,ip:%v", key, requestIp))
			redis.HDel(ipKey, requestIp)
			tologin(response, request)
		} else {
			if url != route.ROUTE_admin_tologin { //更新生存时间
				redis.Expire(key, common.SESSION_expire_time*time.Second)
			}
			action(response, request) //执行方法
		}
	}
}

func init() {
	err := db.InitRedis()
	if err != nil {
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
		mp[common.ROOT_PATH_FILED] = path
		mp[common.WEB_SERVER_PATH_FILED] = webServerPath
		isSuccess := db.UpdateBatchHash(common.SYSTEM_CONFIG_HASH_KEY, mp)
		if !isSuccess {
			os.Exit(1)
			return
		}
	}
	err = db.InitMongo()
	if err != nil {
		os.Exit(1)
		return
	}
	register() //注册路由
}
