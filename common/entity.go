package common

import "time"

type MsgCommon struct {
	Id         string    `json:"id"`         //消息id
	Ip         string    `json:"ip"`         //消息ip
	CreateTime time.Time `json:"createTime"` //消息创建时间
}

type Msg struct {
	MsgCommon
	MsgType uint8  `json:"msgType"` //0-视频
	MsgData []byte `json:"msgData"` //数据
}

type Video struct {
	Name  string `json:"name"`  //文件名称
	Size  uint64 `json:"size"`  //文件大小
	Class string `json:"class"` //文件分类
	Data  []byte `json:"data"`  //文件数据
}
