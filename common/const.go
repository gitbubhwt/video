package common

const (
	SERVER_NET = "tcp"
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
