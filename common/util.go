package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"io"
	"crypto/rand"
	"encoding/base64"
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

//获取当前目录
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/

}

//截取字符串
func Substring(str string, begin, length int) string {
	lth := len(str)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(str[begin:end])
}

//字符串去空格和行且变大写
func GetUpperNull(label string) string {
	label = strings.Replace(label, " ", "", -1)  //去掉空格
	label = strings.Replace(label, "\n", "", -1) //去掉换行符
	label = strings.ToUpper(label)               //大写
	return label
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}