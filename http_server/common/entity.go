package common

type Video struct {
	Index      string `json:"index"`
	ImgSrc     string `json:"imgSrc"`     //视频播放首页图片
	Path       string `json:"path"`       //视频路径
	CreateTime int64  `json:"createTime"` //创建时间
}
