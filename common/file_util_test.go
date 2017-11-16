package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*
获取程序运行路径
*/
func GetCurrentDirectory1() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}