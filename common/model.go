package common

import "time"

type Video struct {
	Id         int64        `xorm:"BIGSERIAL" json:"id"`
	Cover      string       `xorm:"varchar(300)"json:"cover"` //视频播放首页图片
	Name       string       `xorm:"varchar(100)"json:"name"`  //视频名称
	Type       string       `xorm:"varchar(3)" json:"type"`   //类型
	CreateTime time.Time    `xorm:"updated"json:"createTime"` //创建时间
	Path       []*VideoPath //视频路径
}

type VideoPath struct {
	Id         int64     `xorm:"BIGSERIAL" json:"id"`
	VideoId    int64     `xorm:"bigint"json:"videoId"`     //视频播放首页图片
	Path       string    `xorm:"varchar(300)"json:"path"`  //视频路径
	OrderNum   int       `xorm:"int(3)" json:"orderNum"`   //排序`
	CreateTime time.Time `xorm:"updated"json:"createTime"` //创建时间
}

type MonVideo struct {
	Id         string         `json:"id"`
	Cover      string         `json:"cover"`                       //视频播放首页图片
	Name       string         `json:"name"`                        //视频名称
	Type       string         `json:"type"`                        //类型
	CreateTime int64          `bson:"createTime"json:"createTime"` //创建时间
	Path       []MonVideoPath //视频路径
}

type MonVideoPath struct {
	Path     string `json:"path"`                    //视频路径
	OrderNum int    `bson:"orderNum"json:"orderNum"` //排序`
}

type MonAdmin struct {
	Id         string `json:"id"`                          //管理员id
	Name       string `json:"name"`                        //名称
	Pwd        string `json:"pwd"`                         //密码
	Sex        uint8  `json:"sex"`                         //性别
	FullName   string `bson:"fullName" json:"fullName"`    //名字
	CreateTime int64  `bson:"createTime"json:"createTime"` //创建时间
}
