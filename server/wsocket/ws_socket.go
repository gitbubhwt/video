package wsocket

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"video/common"
	log "video/logger"
)

type WsSocket struct {
	MaskingKey []byte
	Conn       net.Conn
}

//创建wssocket
func NewWsSocket(conn net.Conn) *WsSocket {
	return &WsSocket{Conn: conn}
}

//写入数据
func (this *WsSocket) Write(data []byte) error {
	// 这里只处理data长度<125的
	if len(data) >= 125 {
		return errors.New("send ws socket data error")
	}
	lenth := len(data)
	maskedData := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		if this.MaskingKey != nil {
			maskedData[i] = data[i] ^ this.MaskingKey[i%4]
		} else {
			maskedData[i] = data[i]
		}
	}
	this.Conn.Write([]byte{0x81})
	var payLenByte byte
	if this.MaskingKey != nil && len(this.MaskingKey) != 4 {
		payLenByte = byte(0x80) | byte(lenth)
		this.Conn.Write([]byte{payLenByte})
		this.Conn.Write(this.MaskingKey)
	} else {
		payLenByte = byte(0x00) | byte(lenth)
		this.Conn.Write([]byte{payLenByte})
	}
	this.Conn.Write(data)
	return nil
}

//读取数据
func (this *WsSocket) Read() (data []byte, err error) {
	err = nil
	//第一个字节：FIN + RSV1-3 + OPCODE
	opcodeByte := make([]byte, 1)
	this.Conn.Read(opcodeByte)
	FIN := opcodeByte[0] >> 7
	RSV1 := opcodeByte[0] >> 6 & 1
	RSV2 := opcodeByte[0] >> 5 & 1
	RSV3 := opcodeByte[0] >> 4 & 1
	//OPCODE := opcodeByte[0] & 15
	if RSV1 != 0x00 || RSV2 != 0x00 || RSV3 != 0x00 {
		log.Error("Ws read data is wrong")
		return
	}
	//log.Info(RSV1, RSV2, RSV3, OPCODE)
	payloadLenByte := make([]byte, 1)
	this.Conn.Read(payloadLenByte)
	payloadLen := int(payloadLenByte[0] & 0x7F)
	mask := payloadLenByte[0] >> 7
	if payloadLen == 127 {
		extendedByte := make([]byte, 8)
		this.Conn.Read(extendedByte)
	}
	maskingByte := make([]byte, 4)
	if mask == 1 {
		this.Conn.Read(maskingByte)
		this.MaskingKey = maskingByte
	}
	payloadDataByte := make([]byte, payloadLen)
	this.Conn.Read(payloadDataByte)
	//log.Info("data:", payloadDataByte)
	dataByte := make([]byte, payloadLen)
	for i := 0; i < payloadLen; i++ {
		if mask == 1 {
			dataByte[i] = payloadDataByte[i] ^ maskingByte[i%4]
		} else {
			dataByte[i] = payloadDataByte[i]
		}
	}
	if FIN == 1 {
		data = dataByte
		return
	}
	nextData, err := this.Read()
	if err != nil {
		return
	}
	data = append(data, nextData...)
	return
}

//握手
func (this *WsSocket) HandShake(data []byte) bool {
	//解析握手包
	headers := this.parseHandshake(string(data))
	//组装回复包 握手连接
	secWebsocketKey := headers[common.WS_HEADERS_KEY]
	//base64 加密
	h := sha1.New()
	io.WriteString(h, secWebsocketKey+common.WS_QUID)
	accept := make([]byte, 28)
	base64.StdEncoding.Encode(accept, h.Sum(nil))
	response := fmt.Sprintf(common.WS_RESPONSE, string(accept))
	if _, err := this.Conn.Write([]byte(response)); err != nil {
		log.Error("Ws hand shake write data fail,err:", err)
		return false
	}
	return true
}

//解析网页连接的握手包
func (this *WsSocket) parseHandshake(content string) map[string]string {
	headers := make(map[string]string, 10)
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		if len(line) >= 0 {
			words := strings.Split(line, ":")
			if len(words) == 2 {
				headers[strings.Trim(words[0], " ")] = strings.Trim(words[1], " ")

			}
		}
	}
	return headers
}
