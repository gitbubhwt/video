package route

const (
	VIDEO_FILE_ROOT_PATH                = "/upload"
	HTML_ROOT_PATH                      = "/html"
	ROUTE_INDEX_REQUEST                 = "/video"
	ROUTE_PLAY_REQUEST                  = "/play"
	ROUTE_ADMIN_REQUEST                 = "/admin"
	ROUTE_ADMIN_VIDEO_ADD_REQUEST       = ROUTE_ADMIN_REQUEST + "/video/add"
	ROUTE_ADMIN_VIDEO_LIST_REQUEST      = ROUTE_ADMIN_REQUEST + "/video/list"
	ROUTE_ADMIN_VIDEO_UPLOAD_REQUEST    = ROUTE_ADMIN_REQUEST + "/video/upload"
	ROUTE_VIDEO_DEL_REQUEST             = ROUTE_ADMIN_REQUEST + "/video/del"
	ROUTE_ADMIN_VIDEO_SAVE_REQUEST      = ROUTE_ADMIN_REQUEST + "/video/save"
	ROUTE_ADMIN_VIDEO_LIST_DATA_REQUEST = ROUTE_ADMIN_REQUEST + "/video/pageList"
	ROUTE_ADMIN_LOGIN_REQUEST           = ROUTE_ADMIN_REQUEST + "/login"
	ROUTE_ADMIN_INDEX_REQUEST           = ROUTE_ADMIN_REQUEST + "/index"

	ROUTE_FILTER           = "/"
	ROUTE_INDEX_HTML       = "/index.html"
	ROUTE_PLAY_HTML        = "/video.html"
	ROUTE_ADMIN_ADD_HTML   = "/admin/video/video_add.html"
	ROUTE_ADMIN_LIST_HTML  = "/admin/video/video_list.html"
	ROUTE_ADMIN_LOGIN_HTML = "/admin/login/login.html" //登陆页面
	ROUTE_ADMIN_INDEX_HTML = "/admin/index/index.html"       //后台管理首页
)
