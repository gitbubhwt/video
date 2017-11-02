package common

import (
	"bufio"
	"net"
	"testing"
	//"bytes"
	log "video/logger"
)

func Test(t *testing.T) {
	log.Info("单元测试")
	addr := "192.168.96.131:2345"
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Error(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go HandConn(conn)
	}

}
func HandConn(conn net.Conn) {
	//var buf1 bytes.Buffer
	//buf1:=bytes.NewBuffer(conn)
	buf := bufio.NewReader(conn)
	bytes := make([]byte, 1024)
	for {
		n,err:=buf.Read(bytes)
		log.Info(n,err)
		log.Info(string(bytes))
	}

}
