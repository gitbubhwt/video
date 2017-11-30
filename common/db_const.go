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

const (
	TABLE_VIDEO      = "video"
	TABLE_VIDEO_PATH = "video_path"

	DEFAULT_PAGE_SIZE = 10
	LIMIT_SQL         = " limit %s,%d"
	DEFAULT_WHERE_SQL = " "

	GET_TOTAL_COUNT_SQL   = "select count(*) as total_count %s"
	VIDEO_PAGE_LIST_SQL   = "select * from " + TABLE_VIDEO + "%s" + LIMIT_SQL
	VIDEO_PAGE_SQL        = "select * from " + TABLE_VIDEO_PATH + " where video_id=%v and order_num=%v "
	VIDEO_PATH_INNERT_SQL = "insert into " + TABLE_VIDEO_PATH + "(video_id,path,order_num,create_time) values(?,?,?,now())"
)
