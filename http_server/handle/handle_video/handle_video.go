package handle_video

import (
	"net/http"
	webCommon "video/http_server/common"
	"video/http_server/route"
)

//视频播放页面
func VideoHeadHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_HEAD_HTML_PATH)
}
