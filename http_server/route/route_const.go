package route

const (
	VIDEO_FILE_ROOT_PATH = "/upload"
	HTML_ROOT_PATH       = "/html"
	//url路由
	ROUTE_index                = "/video"
	ROUTE_play                 = "/play"
	ROUTE_admin                = "/admin"
	ROUTE_admin_video_add      = ROUTE_admin + "/video/add"
	ROUTE_admin_video_list     = ROUTE_admin + "/video/list"
	ROUTE_admin_video_upload   = ROUTE_admin + "/video/upload"
	ROUTE_admin_video_del      = ROUTE_admin + "/video/del"
	ROUTE_admin_video_save     = ROUTE_admin + "/video/save"
	ROUTE_admin_video_pageList = ROUTE_admin + "/video/pageList"
	ROUTE_admin_tologin        = ROUTE_admin + "/tologin"
	ROUTE_admin_login          = ROUTE_admin + "/login"
	ROUTE_admin_index          = ROUTE_admin + "/index"

	//html页面
	ROUTE_admin_html            = "/index.html"
	ROUTE_play_html             = "/video.html"
	ROUTE_admin_video_add_html  = "/admin/video/video_add.html"
	ROUTE_admin_video_list_html = "/admin/video/video_list.html"
	ROUTE_admin_login_html      = "/admin/login/login.html" //登陆页面
	ROUTE_admin_index_html      = "/admin/index/index.html" //后台管理首页
)
