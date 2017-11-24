package route

const (
	VIDEO_FILE_ROOT_PATH       = "/upload"
	HTML_ROOT_PATH             = "/html"
	ROUTE_INDEX_REQUEST        = "/video"
	ROUTE_PLAY_REQUEST         = "/play"
	ROUTE_ADMIN_REQUEST        = "/admin"
	ROUTE_VIDEO_ADD_REQUEST    = ROUTE_ADMIN_REQUEST + "/video/add"
	ROUTE_VIDEO_LIST_REQUEST   = ROUTE_ADMIN_REQUEST + "/video/list"
	ROUTE_VIDEO_UPLOAD_REQUEST = ROUTE_ADMIN_REQUEST + "/video/upload"
	ROUTE_VIDEO_SAVE_REQUEST   = ROUTE_ADMIN_REQUEST + "/video/save"
	ROUTE_INDEX_HTML_PATH      = "/index.html"
	ROUTE_PLAY_HTML_PATH       = "/video.html"
	ROUTE_ADD_HTML_PATH        = "/admin/video/video_add.html"
	ROUTE_LIST_HTML_PATH       = "/admin/video/video_list.html"
)
