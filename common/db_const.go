package common

//redis 配置
const (
	REDIS_ADDR     = "127.0.0.1:6379"
	REDIS_PASSWORD = ""
	REDIS_database = 5
)

//表配置
const (
	keyPrefix = "video."

	//配置表
	SYSTEM_CONFIG_KEY       = keyPrefix + "config"
	SYSTEM_CONFIG_ROOT_PATH = "root_path" //系统配置路径
)
