package handle_video

import (
	"html/template"
	"net/http"
	"video/http_server/route"
	log "video/logger"
)

func VideoHeadHtml(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(route.ROUTE_HEAD_HTML_PATH)
	if err == nil {
		//w.Write()
		t.Execute(w, nil)
	}
	log.Info(err)
}
