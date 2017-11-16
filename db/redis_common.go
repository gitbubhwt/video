package db

import (
	log "video/logger"
)

//保存或者更新数据 hash 表
func UpdateHash(key, filedName string, value interface{}) bool {
	db := GetClient()
	if err := db.HSet(key, filedName, value).Err(); err != nil {
		log.Error("Update fail,key:", key, "filedName:", filedName, "value:", value, "err:", err)
		return false
	}
	return true
}

//获取表属性数据 hash 表
func GetValue(key, filedName string) interface{} {
	db := GetClient()
	var err error
	if result, err := db.HGet(key, filedName).Result(); err == nil {
		return result
	}
	log.Error("Get value fail,key:", key, "filedName:", filedName, "err:", err)
	return nil
}
