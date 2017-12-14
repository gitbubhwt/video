package handle_login

import (
	"net/http"
	webCommon "video/http_server/common"
	"video/http_server/route"
	//log "video/logger"
)
//视频播放页面
func LoginHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_LOGIN_HTML_PATH, nil)
}
