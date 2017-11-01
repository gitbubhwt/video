package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

type Util struct{}

func (util *Util) Uint642Bytes(data uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, data)
	return b
}

func (util *Util) Bytes2Uint64(data []byte) uint64 {
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
