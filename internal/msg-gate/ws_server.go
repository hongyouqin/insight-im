package msggate

import (
	"bytes"
	"encoding/gob"
	"insight/pkg/common/config"
	"insight/pkg/utils"
	"net/http"
	"strconv"
	"time"

	"insight/pkg/common/constant"

	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func NewWsServer(cfg *config.GateConfig, log *zap.Logger, validate *validator.Validate) *WsServer {
	var w WsServer
	w.cfg = cfg
	w.wsAddr = ":" + cfg.WsSvrCfg.Port
	w.wsMaxConnNum = cfg.WsSvrCfg.MaxConnNum
	w.userConnManager.onInit(log)
	w.log = log
	return &w
}

type WsServer struct {
	wsAddr          string
	wsMaxConnNum    int
	upgrader        *websocket.Upgrader
	userConnManager UserConnManager
	cfg             *config.GateConfig
	log             *zap.Logger
	validate        *validator.Validate
}

func (w *WsServer) StartWs() {
	w.validate = w.validate
	w.upgrader = &websocket.Upgrader{
		HandshakeTimeout: time.Duration(w.cfg.WsSvrCfg.Timeout) * time.Second,
		ReadBufferSize:   w.cfg.WsSvrCfg.MaxMsgLen,
		CheckOrigin:      func(r *http.Request) bool { return true },
	}

	w.log.Info("ws server listen success", zap.String("address", w.wsAddr))

	http.HandleFunc("/", w.wsHandler)
	err := http.ListenAndServe(w.wsAddr, nil)
	if err != nil {
		panic("Ws listening err:" + err.Error())
	}
}

func (ws *WsServer) wsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	operationID := ""
	if len(query["operationID"]) != 0 {
		operationID = query["operationID"][0]
	} else {
		operationID = utils.OperationIDGenerator()
	}

	if query == nil {
		ws.log.Error("args is empty")
		return
	}
	if len(query["token"]) == 0 || len(query["userId"]) == 0 && len(query["platformID"]) == 0 {
		ws.log.Error("args err ", zap.Any("query", query))
		return
	}
	token := query["token"][0]
	userId := query["userId"][0]
	platformID := utils.StringToInt(query["platformID"][0])

	if isPass := ws.checkAuth(w, r); isPass {
		wsConn, err := ws.upgrader.Upgrade(w, r, nil)
		if err != nil {
			ws.log.Error(err.Error())
			return
		}
		newConn := ws.userConnManager.addUserConn(
			userId,
			platformID,
			wsConn,
			token,
			wsConn.RemoteAddr().String()+"_"+strconv.Itoa(int(utils.GetCurrentTimestampByMill())),
			operationID,
		)
		go ws.readMsg(newConn)
	}

}

// 用户token鉴权
func (ws *WsServer) checkAuth(w http.ResponseWriter, r *http.Request) (isPass bool) {
	//验证token的合法性
	isPass = true
	return
}

func (ws *WsServer) readMsg(conn *Conn) {
	for {
		msgType, msg, err := conn.ws.ReadMessage()
		if msgType == websocket.PingMessage {
			ws.log.Info("this is a  pong")
		}
		if err != nil {
			ws.log.Error("ws readmsg error", zap.String("error", err.Error()), zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
			ws.userConnManager.delUserConn(conn)
			return
		}
		ws.msgParse(conn, msg)
	}
}

func (ws *WsServer) msgParse(conn *Conn, msg []byte) {
	// to do
	b := bytes.NewBuffer(msg)
	decoder := gob.NewDecoder(b)
	input := Req{}
	err := decoder.Decode(&input)
	if err != nil {
		ws.log.Error("msg parse error", zap.String("error", err.Error()), zap.String("error", err.Error()), zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
		return
	}

	if err := ws.validate.Struct(input); err != nil {
		ws.log.Error("", zap.String("err", err.Error()))
		//ws.sendErrMsg(conn, 201, err.Error(), m.ReqIdentifier, m.MsgIncr, m.OperationID)
		return
	}

	if input.SendID != conn.userId {
		if err = conn.ws.Close(); err != nil {
			ws.log.Error("close ws conn failed", zap.String("error", err.Error()), zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
			return
		}
		ws.log.Error("close ws conn", zap.String("error", err.Error()), zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
		return
	}

	switch input.ReqIdentifier {
	case constant.WSSendMsg:
		//转发消息给msg服务
		ws.sendMsg(conn, &input)
	case constant.WSHeartbeat:
		//这里的心跳，赋予新的功能，会用于消息的同步处理
		ws.heartbeat(conn, &input)
	default:
		ws.log.Error("ReqIdentifier failed ", zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
	}

}

func (ws *WsServer) heartbeat(conn *Conn, msgReq *Req) {
	// resp := api.MsgResp{
	// 	Type:      api.PackageType_PT_HEARTBEAT,
	// 	RequestId: msgReq.RequestId,
	// 	Code:      0,
	// }
	// msg, err := proto.Marshal(&resp)
	// if err != nil {
	// 	ws.log.Error("heartbeat msg error", zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
	// 	return
	// }

	// conn.wsMutex.Lock()
	// defer conn.wsMutex.Unlock()
	// conn.ws.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	// err = conn.ws.WriteMessage(websocket.BinaryMessage, msg)
	// if err != nil {
	// 	ws.log.Error("send Heartbeat ws writermsg error", zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
	// }
}

// 转发消息
func (ws *WsServer) sendMsg(conn *Conn, msg *Req) {
	//to do
	ws.log.Info("消息投递", zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
}
