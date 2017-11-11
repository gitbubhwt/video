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

var logOpen bool = false
var logChannel chan interface{}

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
	//log.Println(v)
	//openLog()
	//writeLog(DEBUG, caller(), v)
	log.Println(INFO, caller(), v)
}

func Error(v ...interface{}) {
	//openLog()
	//writeLog(DEBUG, caller(), v)
	log.Println(ERROR, caller(), v)
}
func Debug(v ...interface{}) {
	//openLog()
	//writeLog(DEBUG, caller(), v)
	log.Println(DEBUG, caller(), v)
}

func openLog() {
	if logOpen == false {
		logOpen = true
		go func() {
			logChannel = make(chan interface{})
			for {
				v ,ok:= <-logChannel
				if !ok{
					break
				}
				log.Println(v)
			}
			logOpen = false
		}()
	}
}

func writeLog(v ...interface{}) {
	logChannel <- v
}
