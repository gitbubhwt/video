package common

import (
	"encoding/json"
	"reflect"
)

type MsgCommon struct {
	Id string `json:"id"` //消息id
	Ip string `json:"ip"` //消息ip
}

func (this MsgCommon) String() string {
	obj := reflect.ValueOf(&this)
	return StructPrint(obj)
}

type Msg struct {
	From       MsgCommon   `json:"from"`
	To         MsgCommon   `json:"to"`
	CreateTime int64       `json:"createTime"` //消息创建时间
	MsgType    uint8       `json:"msgType"`    //0-心跳  1-视频
	Content    interface{} `json:"content"`    //数据
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

func (this VideoClient) String() string {
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

func (this VideoServer) String() string {
	obj := reflect.ValueOf(&this)
	return StructPrint(obj)
}

//当前播放视频状态
type VideoState struct {
	Name        string `json:"name"`        //视频名称
	State       uint8  `json:"state"`       //视频状态
	CurrentTime int64  `json:"currentTime"` //当前播放视频时间
}

type Ack struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (ack *Ack) ResponseError() []byte {
	ack.Code = ACK_FAIL
	data, _ := json.Marshal(ack)
	return data
}
func (ack *Ack) ResponseSuccess() []byte {
	ack.Code = ACK_SUCCESS
	data, _ := json.Marshal(ack)
	return data
}
