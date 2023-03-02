package msggate

import (
	"insight/pkg/common/constant"
	"insight/pkg/utils"
	"sync"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// 管理用户的conn链接
type UserConnManager struct {
	rwLock        *sync.RWMutex
	log           *zap.Logger
	userConnCount uint64
	wsConnToUser  map[*Conn]map[int]string //用户id和conn连接的对应关系 支持多端 一对一
	wsUserToConn  map[string]map[int]*Conn //用户id和conn连接的对应关系 支持多端 一对多
}

func (uc *UserConnManager) onInit(log *zap.Logger) {
	uc.rwLock = new(sync.RWMutex)
	uc.log = log
	uc.userConnCount = 0
	uc.wsConnToUser = make(map[*Conn]map[int]string)
	uc.wsUserToConn = make(map[string]map[int]*Conn)
}

func (uc *UserConnManager) addUserConn(uid string, platformID int, conn *websocket.Conn, token string, connID, operationID string) (newConn *Conn) {
	uc.rwLock.Lock()
	defer uc.rwLock.Unlock()

	newConn = &Conn{
		connType:   ConnModeWs,
		ws:         conn,
		userId:     uid,
		token:      token,
		PlatformID: platformID,
		connID:     connID,
	}
	uc.log.Info("add user conn",
		zap.String("func: ", utils.GetSelfFuncName()),
		zap.String("uid", uid),
		zap.String("top", token),
		zap.String("ip", conn.RemoteAddr().String()))

	//user to conn 的映射map 一对多
	if oldConnMap, ok := uc.wsUserToConn[uid]; ok {
		oldConnMap[platformID] = newConn
		uc.wsUserToConn[uid] = oldConnMap
		uc.log.Sugar().Info("user not first come in, add conn ", uid, platformID, conn, oldConnMap)
	} else {
		i := make(map[int]*Conn)
		i[platformID] = newConn
		uc.wsUserToConn[uid] = i
		uc.log.Sugar().Info("user first come in, add conn ", uid, platformID, conn, oldConnMap)
	}

	//conn to user 的映射map 一对一
	if oldStringMap, ok := uc.wsConnToUser[newConn]; ok {
		oldStringMap[platformID] = uid
		uc.wsConnToUser[newConn] = oldStringMap
	} else {
		i := make(map[int]string)
		i[platformID] = uid
		uc.wsConnToUser[newConn] = i
	}

	uc.userConnCount++
	count := 0
	for _, v := range uc.wsUserToConn {
		count = count + len(v)
	}
	uc.log.Sugar().Info("wsUser added", "connection_uid", uid, "connection_platform", constant.PlatformIDToName(int32(platformID)), "online_user_num", len(uc.wsUserToConn), "online_conn_num", count)
	return
}

func (uc *UserConnManager) delUserConn(conn *Conn) {
	uc.rwLock.Lock()
	defer uc.rwLock.Unlock()
	uc.userConnCount--
	var uid string
	var platform int
	if oldStringMap, okg := uc.wsConnToUser[conn]; okg {
		for k, v := range oldStringMap {
			platform = k
			uid = v
		}
		if oldConnMap, ok := uc.wsUserToConn[uid]; ok {
			delete(oldConnMap, platform)
			uc.wsUserToConn[uid] = oldConnMap
			if len(oldConnMap) == 0 { //用户最后一个链接
				delete(uc.wsUserToConn, uid)
			}
			count := 0
			for _, v := range uc.wsUserToConn {
				count = count + len(v)
			}
			uc.log.Sugar().Info("WS delete operation", "", "wsUser deleted", uc.wsUserToConn, "disconnection_uid", uid, "disconnection_platform", platform, "online_user_num", len(uc.wsUserToConn), "online_conn_num", count)
		}
		delete(uc.wsConnToUser, conn)
	}
	err := conn.ws.Close()
	if err != nil {
		uc.log.Sugar().Error(" close err", "", "uid", uid, "platform", platform)
	}
	if conn.PlatformID == 0 || conn.connID == "" {
		uc.log.Sugar().Warn(utils.GetSelfFuncName(), "PlatformID or connID is null", conn.PlatformID, conn.connID)
	}
}

func (uc *UserConnManager) getUserConn(uid string, platform int) *Conn {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()
	if connMap, ok := uc.wsUserToConn[uid]; ok {
		if conn, flag := connMap[platform]; flag {
			return conn
		}
	}
	return nil
}

func (uc *UserConnManager) getUserAllCons(uid string) map[int]*Conn {
	uc.rwLock.RLock()
	defer uc.rwLock.RUnlock()
	if connMap, ok := uc.wsUserToConn[uid]; ok {
		newConnMap := make(map[int]*Conn)
		for k, v := range connMap {
			newConnMap[k] = v
		}
		return newConnMap
	}
	return nil
}
