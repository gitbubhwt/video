package common

import (
	"html/template"
	"net/http"
	"strings"
	"video/common"
	log "video/logger"
)

//跳转页面
func GoToPage(w http.ResponseWriter, htmlPath string, data interface{}) {
	directory := common.GetCurrentDirectory()
	log.Info(directory)
	if strings.Index(directory, SERVER_ROOT_PATH) != -1 {
		htmlPath = directory + htmlPath
	} else {
		htmlPath = SERVER_ROOT_PATH + htmlPath
	}
	if t, err := template.ParseFiles(htmlPath); err == nil {
		t.Execute(w, data)
	} else {
		log.Error(err)
	}
}
