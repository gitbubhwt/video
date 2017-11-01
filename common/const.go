package common

const (
	SERVER_NET      = "tcp"
	LOG_HEAD_SERVER = "Server"
	LOG_HEAD_CLIENT = "Client"
	HEART_RATE      = 10 //秒
	//READ_ERROR_COUNT = 3  //允许包错误3次,连续3次错误，断开连接
	UPLOAD_COMPLETE       = 1 //上传完成
	UPLOAD_CONTINUE       = 0 //继续上传
	CLIENT_FILE_ROOT_PATH = "C:/icsoc/"
)
const (
	PACKAGE_HEAD_A = 0xaa
	PACKAGE_HEAD_B = 0xbb
	PACKAGE_END    = 0xcc
)

const (
	MessageType_MSG_TYPE_HEART = 0 //心跳
	MessageType_MSG_TYPE_VEDIO = 1 //视频
)

const (
	SERVER_TYPE_SOCKET  = 0 //socket
	SERVER_TYPE_WSOCKET = 1 //wSocket服务
)
