package common

import "time"

type Video struct {
	Index      string    `xorm:"BIGSERIAL" json:"index"`
	ImgSrc     string    `xorm:"varchar(300)"json:"imgSrc"` //视频播放首页图片
	Name       string    `xorm:"varchar(100)"json:"name"`   //视频名称
	Path       string    `xorm:"varchar(300)"json:"path"`   //视频路径
	CreateTime time.Time `xorm:"updated"json:"createTime"`  //创建时间
}
