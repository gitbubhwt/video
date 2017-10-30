package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

const (
	INFO  = "[INFO]"
	ERROR = "[ERROR]"
	DEBUG = "[DEBUG]"
)

// 获取调用文件和行号
func caller() string {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		line = 0
	}
	pathArgs := strings.Split(file, "/")
	filename := pathArgs[len(pathArgs)-1]
	packAndFuncName := runtime.FuncForPC(pc).Name()
	packFunc := strings.Split(packAndFuncName, ".")
	funcName := packFunc[len(packFunc)-1]
	packName := strings.Join(packFunc[:len(packFunc)-1], ".")
	return fmt.Sprintf("(%s %s %s:%d)", packName, filename, funcName, line)
}

func Info(v ...interface{}) {
	log.Println(INFO, caller(), v)
}

func Error(v ...interface{}) {
	log.Println(ERROR, caller(), v)
}
func Debug(v ...interface{}) {
	log.Println(DEBUG, caller(), v)
}
