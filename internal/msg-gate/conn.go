package msggate

import (
	"net"
	"sync"

	"github.com/gorilla/websocket"
)

const (
	//连接方式
	ConnModeTcp = 1 //tcp模式
	ConnModeWs  = 2 //ws模式
)

type Conn struct {
	connType   int8            // 连接方式
	tc         *net.TCPConn    //TCP连接
	ws         *websocket.Conn //Websocket连接
	wsMutex    sync.Mutex
	userId     string //用户id
	token      string
	PlatformID int    //平台id
	connID     string //连接id
}
