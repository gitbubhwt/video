package main

import (
	"fmt"
	"time"
	"video/db"
)

type Action func(arg string) string

var ma map[string]func(arg string) string

func main() {
	redisClient := db.GetRedisClient()
	key := fmt.Sprintf("testttl")
	redisClient.HSet(key, "test", "1")
	if result, err := redisClient.TTL(key).Result(); err == nil {
		if result == -1 { //key存在但没有设置生存时间
			fmt.Println("not exist")
			redisClient.Expire(key, 60*time.Second)
		} else if result == -2 { //不存在
			fmt.Println("remove")
		} else if result >= 0 { //存在,剩余时间
			fmt.Println("time:", result)
		}
	}
}

func test(arg string) string {
	fmt.Println("test:", arg)
	return fmt.Sprintf("show:%s", arg)
}

func show(rt string) Action {
	if t, ok := ma[rt]; ok {
		return t
	}
	return nil
}
