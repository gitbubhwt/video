package common

import (
	"bufio"
	"encoding/json"
	"io"
	"mime/multipart"
	"net"
	"os"
	"time"
	log "video/logger"
)

//读文件
func ReadFile(path string, conn net.Conn) {
	isExist := checkFileIsExist(path)
	if isExist {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			log.Error("open file fail,path:", path)
			return
		}
		buf := make([]byte, 1024)
		buffer := bufio.NewReader(file)
		msg := new(Msg) //发送消息
		msg.CreateTime = time.Now().Unix()
		msg.MsgType = MessageType_MSG_TYPE_VEDIO
		video := new(VideoServer)
		video.Name = path
		var off int64
		for {
			n, err := buffer.Read(buf)
			if err != nil && err != io.EOF {
				log.Error("read file fail,err:", err)
				break
			}
			if n == 0 {
				log.Info("read file complete,path:", path)
				break
			}
			video.Data = buf[:n]
			video.Off = off
			video.Complete = UPLOAD_CONTINUE
			bytes := PackFileMsg(msg, video)
			SendMsgByTcp(bytes, conn, "Client send file")
			off += int64(n)
		}
		video.Data = nil
		video.Off = 0
		video.Complete = UPLOAD_COMPLETE
		bytes := PackFileMsg(msg, video)
		SendMsgByTcp(bytes, conn, "Client send file")
	}
}

//打包文件消息
func PackFileMsg(msg *Msg, video *VideoServer) []byte {
	data, err := json.Marshal(video)
	if err != nil {
		log.Error("read file fail,err:", err)
		return nil
	}
	msg.Content = data
	bytes, err := json.Marshal(msg)
	if err != nil {
		log.Error("read file fail,err:", err)
		return nil
	}
	return bytes
}

//写文件
func WriteFile(path string, data []byte, off int64) {
	var file *os.File
	var err error
	if !checkFileIsExist(path) {
		file, err = os.Create(path)
	}
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Error("open file fail,path:", path)
		return
	}
	file.WriteAt(data, off)
}

//创建文件
func CreateFile(path string, uploadFile multipart.File) error {
	var file *os.File
	var err error
	if !checkFileIsExist(path) {
		_, err = os.Create(path)
		if err != nil {
			return err
		}
	}
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	_, err = io.Copy(file, uploadFile)
	if err != nil {
		return err
	}
	defer file.Close()
	defer uploadFile.Close()
	return nil
}

//判断文件是否存在
func checkFileIsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
