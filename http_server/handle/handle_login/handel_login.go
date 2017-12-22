package handle_login

import (
	"net/http"
	webCommon "video/http_server/common"
	"video/http_server/route"
	//log "video/logger"
)
//视频播放页面
func AdminLoginHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_ADMIN_LOGIN_HTML, nil)
}
