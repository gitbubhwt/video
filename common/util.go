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
	//s, err := exec.LookPath(os.Args[0])
	//checkErr(err)
	//i := strings.LastIndex(s, "\\")
	//path := string(s[0 : i+1])
	//path = strings.Replace(path, "\\", "/", -1)
	//return path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
