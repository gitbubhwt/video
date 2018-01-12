package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"video/common"
	"video/db"
	log "video/logger"
)

//跳转页面
func GoToPage(w http.ResponseWriter, htmlPath string, data interface{}) {
	rootPath := db.GetStringValue(common.SYSTEM_CONFIG_HASH_KEY, common.WEB_SERVER_PATH_FILED)
	htmlPath = rootPath + WEB_SERVER_HTML_PATH + htmlPath
	if t, err := template.ParseFiles(htmlPath); err == nil {
		if err := t.Execute(w, data); err == nil {
			log.Info(fmt.Sprintf("Go to page success,htmlPath:%v", htmlPath))
		} else {
			log.Error(fmt.Sprintf("Go to page fail,htmlPath:%v,err:%v", htmlPath, err))
		}
	} else {
		log.Error(fmt.Sprintf("Go to page fail,err:%v", err))
	}
}

//提示响应
func GoToResponse(w http.ResponseWriter, code int, msg string, data interface{}) {
	ack := new(common.Ack)
	ack.Msg = msg
	ack.Code = code
	ack.Data = data
	if data, err := json.Marshal(ack); err == nil {
		w.Write(data)
	} else {
		log.Error("Go to response fail,err:", err)
	}
}

//发送响应
func SendResponse(w http.ResponseWriter, data interface{}) {
	if data, err := json.Marshal(data); err == nil {
		w.Write(data)
	} else {
		log.Error("Send response fail,err:", err)
	}
}

//分页
func GetPageOption(pageNo string, pageSize int64, sql string) (*PageOption, error) {
	//sql select * from user where user_name like='dddd'
	//获取数据总数
	end := strings.Index(sql, "from")
	getTotalCountSql := fmt.Sprintf(common.GET_TOTAL_COUNT_SQL, common.Substring(sql, end, len(sql)))
	sqlDb := db.GetMysql()
	results, err := sqlDb.QueryString(getTotalCountSql)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("get total count fail,results size is zero")
	}
	var totalCount int64
	totalCount, err = strconv.ParseInt(results[0]["total_count"], 10, 64)
	if err != nil {
		return nil, err
	}
	pageOption := new(PageOption)
	pageOption.PageNo = pageNo
	pageOption.PageSize = pageSize
	pageOption.TotalCount = totalCount
	//计算总页数
	var totalPage int64
	if totalCount%pageSize == 0 {
		totalPage = totalCount / pageSize
	} else {
		totalPage = totalCount/pageSize + 1
	}
	pageOption.TotalPage = totalPage
	return pageOption, nil
}

type MongoPageOption struct {
	IsHome       bool        `json:"isHome"`
	IsEnd        bool        `json:"isEnd"`
	IsShow       bool        `json:"isShow"`
	PageNo       int         `json:"pageNo"`
	PageSize     int         `json:"pageSize"`
	TotalPage    int         `json:"totalPage"`
	TotalCount   int         `json:"totalCount"`
	List         interface{} `json:"list"`
	PageText     string      `json:"pageText"`
	PageSizeText string      `json:"pageSizeText"`
}

//分页
func (pageOption *MongoPageOption) GetMongoPageOption(query *mgo.Query, data interface{}) (*MongoPageOption, error) {
	totalCount, err := query.Count()
	if err != nil {
		return nil, err
	}
	pageSize := pageOption.PageSize
	pageOption.TotalCount = totalCount
	//计算总页数
	var totalPage int
	if totalCount%pageSize == 0 {
		totalPage = totalCount / pageSize
	} else {
		totalPage = totalCount/pageSize + 1
	}
	pageOption.IsShow = true
	pageOption.TotalPage = totalPage
	pageOption.List = data
	if pageOption.TotalPage <= 1 {
		pageOption.IsShow = false
		return pageOption, nil
	}
	pageOption.IsHome = false
	pageOption.IsEnd = false
	if pageOption.PageNo <= 1 {
		pageOption.IsHome = true
	}
	if pageOption.PageNo >= pageOption.TotalPage {
		pageOption.IsEnd = true
	}
	startCount := (pageOption.PageNo-1)*pageOption.PageSize + 1
	endCount := pageOption.PageNo * pageOption.PageSize
	if endCount > pageOption.TotalCount {
		endCount = pageOption.TotalCount
	}
	pageOption.PageText = fmt.Sprintf(common.PAGE_TEXT, startCount, endCount, pageOption.TotalCount)
	pageOption.PageSizeText = fmt.Sprintf(common.PAGE_SIZE_TEXT, pageOption.PageNo, pageOption.TotalPage)
	return pageOption, nil
}
