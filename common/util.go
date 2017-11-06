package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func Uint642Bytes(data uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, data)
	return b
}

func Bytes2Uint64(data []byte) uint64 {
	var t uint64
	buf := bytes.NewBuffer(data)
	binary.Read(buf, binary.BigEndian, &t)
	return t
}

func StructPrint(obj reflect.Value) string {
	var buf bytes.Buffer
	var log string
	ref := obj.Elem()
	nameTypes := ref.Type()
	for i := 0; i < ref.NumField(); i++ {
		filed := ref.Field(i)
		log = fmt.Sprintf("%s:%v", nameTypes.Field(i).Name, filed.Interface())
		buf.WriteString(log)
		buf.WriteString(" ")
	}
	return string(buf.Bytes())
}

//字符串截取
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取父目录
func GetParentDirectory(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//获取当前目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
