package common

//表配置
const (
	keyPrefix = "video."
	//配置表
	SYSTEM_CONFIG_HASH_KEY = keyPrefix + "config"
	WEB_SERVER_PATH_FILED  = "web_server_path" //web系统服务路径
	ROOT_PATH_FILED        = "root_path"       //系统服务路径
	//会话表
	SESSION_HASH_KEY      = keyPrefix + "session.%s.%v"
	SESSION_F_USER_NAME   = "user_name"
	SESSION_F_PWD         = "pwd"
	SESSION_F_CREATE_TIME = "create_time"
	//ip对应的会话sid
	IP_SESSION_HASH_KEY = keyPrefix + "ip.session.%s"
)

const (
	DB_user             = "user"
	DB_admin            = "admin"
	SESSION_expire_time = 60 //秒

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

const (
	MONGO_COLLECTION_VIDEO = "video"
)
