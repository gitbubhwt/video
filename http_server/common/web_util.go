package common

import (
	"html/template"
	"net/http"
	"video/common"
	"video/db"
	log "video/logger"
)

//跳转页面
func GoToPage(w http.ResponseWriter, htmlPath string, data interface{}) {
	rootPathT := db.GetValue(common.SYSTEM_CONFIG_KEY, common.SYSTEM_CONFIG_ROOT_PATH)
	if rootPath, ok := rootPathT.(string); ok && rootPath != "" {
		htmlPath = rootPath + WEN_SERVER_HTML_PATH + htmlPath
		if t, err := template.ParseFiles(htmlPath); err == nil {
			t.Execute(w, data)
		} else {
			log.Error(err)
		}
	} else {
		log.Error(common.SYSTEM_CONFIG_ROOT_PATH, "type is wrong")
	}
}
