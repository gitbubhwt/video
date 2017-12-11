package common

type VideoPlay struct {
	VideoId string //视频id
	Cover   string //封面路径
	Path    string //路径
	Order   string //序号
}

type PageOption struct {
	PageNo     string       `json:"pageNo"`
	PageSize   int64       `json:"pageSize"`
	TotalPage  int64       `json:"totalPage"`
	TotalCount int64       `json:"totalCount"`
	List       interface{} `json:"list"`
}
