package handle_index

import (
	"net/http"
	webCommon "video/http_server/common"
	"video/http_server/route"
	//log "video/logger"
)

//视频播放页面
func AdminIndexHtml(w http.ResponseWriter, r *http.Request) {
	webCommon.GoToPage(w, route.ROUTE_ADMIN_INDEX_HTML, nil)
}
