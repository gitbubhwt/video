package common

const (
	SERVER_ROOT_PATH       = "http_server"
	WEB_SERVER_STATIC_PATH = "/static"
	WEB_SERVER_HTML_PATH   = WEB_SERVER_STATIC_PATH + "/html"
	WEB_SERVER_CSS         = "/css/"
	WEB_SERVER_JS          = "/js/"
	WEB_SERVER_IMG         = "/img/"
	WEB_SERVER_UPLOAD      = "/upload/"
)

const (
	HEAD_VIDEO_ID    = "id"
	HEAD_VIDEO_ORDER = "order"
)

const (
	WEB_SERVER_UPLOAD_FILE_TEMP_PATH = "C:/FFOutput/video/upload/%s"
	WEB_SERVER_UPLOAD_FILE_PATH      = WEB_SERVER_STATIC_PATH + "/upload/%s" //文件上传路径
)

const (
	//请求类型
	METHOD_GET  = "GET"
	METHOD_POST = "POST"
	//会话超时结束
	ADMIN_SESSION_time_out = "admin_session_time_out"
	//登录参数
	ADMIN_LOGIN_PARAM_userName = "admin_userName"
	ADMIN_LOGIN_PARAM_pwd      = "admin_pwd"
)
