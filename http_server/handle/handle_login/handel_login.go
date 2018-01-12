package handle_login

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"video/common"
	"video/db"
	webCommon "video/http_server/common"
	"video/http_server/route"
	log "video/logger"
)

//后台登录页面
func AdminToLoginHtml(response http.ResponseWriter, request *http.Request) {
	log.Info("Admin to login,method:%v", request.Method)
	switch request.Method {
	case webCommon.METHOD_GET:
		{
			webCommon.GoToPage(response, route.ROUTE_admin_login_html, nil)
		}
	case webCommon.METHOD_POST:
		{
			webCommon.GoToResponse(response, common.ACK_TIME_OUT, webCommon.ADMIN_SESSION_time_out,nil)
		}
	}
}

//后台登录
func AdminLogin(w http.ResponseWriter, request *http.Request) {
	userName := request.FormValue(webCommon.ADMIN_LOGIN_PARAM_userName) //userName
	pwd := request.FormValue(webCommon.ADMIN_LOGIN_PARAM_pwd)           //pwd
	log.Info(fmt.Sprintf("Admin login,input params userName:%v,pwd:%v", userName, pwd))
	var requestIp string
	if strings.Index(request.RemoteAddr, ":") != -1 {
		arr := strings.Split(request.RemoteAddr, ":")
		requestIp = arr[0]
	}
	msg := fmt.Sprintf("用户名或者密码错误")
	if userName == "admin" && pwd == "admin" {
		if err := createSession(userName, pwd, requestIp); err == nil {
			msg = fmt.Sprintf("Admin login success")
			log.Error(msg)
			webCommon.GoToResponse(w, common.ACK_SUCCESS, msg,nil)
			return
		} else {
			msg = fmt.Sprintf("Admin login fail,%v", err)
		}
	}
	log.Error(msg)
	webCommon.GoToResponse(w, common.ACK_FAIL, msg,nil)
}

//创建会话
func createSession(userName, pwd, requestIp string) error {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err == nil {
		sid := base64.URLEncoding.EncodeToString(b)
		key := fmt.Sprintf(common.IP_SESSION_HASH_KEY, common.DB_admin)
		db.UpdateHash(key, requestIp, sid)
		//持久化会话
		mp := make(map[string]interface{})
		key = fmt.Sprintf(common.SESSION_HASH_KEY, common.DB_admin, sid)
		mp[common.SESSION_F_USER_NAME] = userName
		mp[common.SESSION_F_PWD] = pwd
		mp[common.SESSION_F_CREATE_TIME] = time.Now().Unix()
		db.UpdateBatchHash(key, mp)
		db.GetRedisClient().Expire(key, common.SESSION_expire_time*time.Second) //设置定时时间
		return nil
	} else {
		return errors.New(fmt.Sprintf("Create session fail,err:%v", err))
	}
}
