package common

import "time"

type Video struct {
	Id         int64     `xorm:"BIGSERIAL" json:"id"`
	Cover      string    `xorm:"varchar(300)"json:"cover"` //视频播放首页图片
	Name       string    `xorm:"varchar(100)"json:"name"`  //视频名称
	Type       string    `xorm:"varchar(3)" json:"type"`   //类型
	CreateTime time.Time `xorm:"updated"json:"createTime"` //创建时间
}

type VideoPath struct {
	Id         string    `xorm:"BIGSERIAL" json:"id"`
	VideoId    int64     `xorm:"int(11)"json:"videoId"`     //视频播放首页图片
	Path       string    `xorm:"varchar(300)"json:"path"`  //视频路径
	Order      int       `xorm:"int(3)" json:"order"`      //排序`
	CreateTime time.Time `xorm:"updated"json:"createTime"` //创建时间
}
