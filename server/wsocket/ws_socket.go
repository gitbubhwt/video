package wsocket

import (
	"crypto/sha1"
	"encoding/base64"
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
	OpcodeByte []byte
}

//创建wssocket
func NewWsSocket(conn net.Conn) *WsSocket {
	return &WsSocket{Conn: conn}
}

//写入数据
func (this *WsSocket) Write(msgByte []byte) error {
	// 掩码开始位置
	masking_key_startIndex  := 2
	var length int = len(msgByte)
	// 计算掩码开始位置
	if length <= 125 {
		masking_key_startIndex = 2
	} else if length > 65536 {
		masking_key_startIndex = 10
	} else if length > 125 {
		masking_key_startIndex = 4
	}
	// 创建返回数据
	result  := make([]byte, masking_key_startIndex+length)
	// 开始计算ws-frame
	// frame-fin + frame-rsv1 + frame-rsv2 + frame-rsv3 + frame-opcode
	result[0] = 0x81 // 129
	// frame-masked+frame-payload-length
	// 从第9个字节开始是 1111101=125,掩码是第3-第6个数据
	// 从第9个字节开始是 1111110>=126,掩码是第5-第8个数据
	if length <= 125 {
		result[1] = byte(length)
	} else if length > 65536 {
		result[1] = 0x7F // 127
	} else if length > 125 {
		result[1] = 0x7E // 126
		result[2] = byte(length >> 8)
		result[3] = byte(length % 256)
	}
	// 将数据编码放到最后
	for i := 0; i < length; i++ {
		result[i+masking_key_startIndex] = msgByte[i]
	}
	if _,err:=this.Conn.Write(result);err!=nil{
		return err
	}
	return nil
}

//读取数据
func (this *WsSocket) Read() (data []byte, err error) {
	err = nil
	//第一个字节：FIN + RSV1-3 + OPCODE
	opcodeByte := this.OpcodeByte
	//this.Conn.Read(opcodeByte)
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
func (this *WsSocket) HandShake() bool {
	buf := make([]byte, 1024)
	n, err := this.Conn.Read(buf)
	if err != nil {
		log.Error("Ws hand shake read data fail,err:", err)
		return false
	}
	data := make([]byte, 0)
	data = append(data, this.OpcodeByte...)
	data = append(data, buf[:n]...)
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
