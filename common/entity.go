package common

import (
	"reflect"
)

type MsgCommon struct {
	Id         string `json:"id"`         //消息id
	Ip         string `json:"ip"`         //消息ip
	CreateTime int64  `json:"createTime"` //消息创建时间
}

func (this MsgCommon) String() string {
	obj := reflect.ValueOf(&this)
	return StructPrint(obj)
}

type Msg struct {
	MsgCommon
	MsgType uint8  `json:"msgType"` //0-心跳  1-视频
	MsgData []byte `json:"msgData"` //数据
}

func (this Msg) String() string {
	obj := reflect.ValueOf(&this)
	return StructPrint(obj)
}

//视频客户端协议
type VideoClient struct {
	Name  string `json:"name"`  //文件名称
	Class string `json:"class"` //文件分类
}

func (this Msg) String() string {
	obj := reflect.ValueOf(&this)
	return StructPrint(obj)
}

//视频服务端协议
type VideoServer struct {
	Name     string `json:"name"`     //文件名称
	Size     uint64 `json:"size"`     //文件大小
	Class    string `json:"class"`    //文件分类
	Data     []byte `json:"data"`     //文件数据
	Off      int64  `json:"off"`      //前后文件标志位
	Complete uint8  `json:"complete"` //是否上传完
}
