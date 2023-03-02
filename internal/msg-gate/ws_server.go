package msggate

import (
	"bytes"
	"context"
	"encoding/gob"
	"insight/pkg/common/config"
	"insight/pkg/utils"
	"net/http"
	"strconv"
	"time"

	rpc "insight/pkg/proto/msg"

	"insight/pkg/common/constant"

	"github.com/go-playground/validator"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
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
		ws.sendMsgReq(conn, &input)
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
func (ws *WsServer) sendMsgReq(conn *Conn, req *Req) {
	//to do
	//ws.log.Info("消息投递", zap.String("userIp", conn.ws.RemoteAddr().String()), zap.String("userId", conn.userId))
	nReplay := new(rpc.SendMsgResp)
	isPass, errCode, errMsg, data := ws.argsValidate(req, constant.WSSendMsg)
	if isPass {
		pdData := rpc.SendMsgReq{
			Token:       req.Token,
			OperationID: req.OperationID,
			Data:        data.(*rpc.MsgData),
		}
		//消息服务grpc客户端,后续用服务发现来替换
		clientConn, err := grpc.Dial("127.0.0.1:7749", grpc.WithInsecure())
		if err != nil {
			ws.log.Error("msg conn failed")
			nReplay.ErrCode = 201
			nReplay.ErrMsg = err.Error()
			ws.sendMsgResp(conn, req, nReplay)
			return
		}
		client := rpc.NewChatClient(clientConn)
		resp, err := client.SendMsg(context.Background(), &pdData)
		if err != nil {
			ws.log.Error("send msg failed", zap.String("err", errMsg))
			nReplay.ErrCode = 200
			nReplay.ErrMsg = err.Error()
			ws.sendMsgResp(conn, req, nReplay)
			return
		}
		ws.sendMsgResp(conn, req, resp)
		ws.log.Info("sendMsgResp rpc call success", zap.String("reply", resp.String()))
	} else {
		nReplay.ErrCode = errCode
		nReplay.ErrMsg = errMsg
		ws.sendMsgResp(conn, req, nReplay)
	}
}

func (ws *WsServer) sendMsgResp(conn *Conn, m *Req, pb *rpc.SendMsgResp) {
	// := make(map[string]interface{})

	var mReplyData rpc.UserSendMsgResp
	mReplyData.ClientMsgID = pb.GetClientMsgID()
	mReplyData.ServerMsgID = pb.GetServerMsgID()
	mReplyData.SendTime = pb.GetSendTime()

	b, _ := proto.Marshal(&mReplyData)
	mReply := Resp{
		ReqIdentifier: m.ReqIdentifier,
		MsgIncr:       m.MsgIncr,
		ErrCode:       pb.GetErrCode(),
		ErrMsg:        pb.GetErrMsg(),
		OperationID:   m.OperationID,
		Data:          b,
	}
	ws.Send(conn, mReply)
}

// 发送答复消息
func (ws *WsServer) Send(conn *Conn, mReply interface{}) {
	var b bytes.Buffer
	//消息序列化
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(mReply)
	if err != nil {
		uid := conn.userId
		platform := conn.PlatformID
		ws.log.Sugar().Error(mReply.(Resp).OperationID, mReply.(Resp).ReqIdentifier, mReply.(Resp).ErrCode, mReply.(Resp).ErrMsg, "Encode Msg error", conn.ws.RemoteAddr().String(), uid, platform, err.Error())
		return
	}
	err = ws.writeMsg(conn, b.Bytes())
	if err != nil {
		uid := conn.userId
		platform := conn.PlatformID
		ws.log.Sugar().Error(mReply.(Resp).OperationID, mReply.(Resp).ReqIdentifier, mReply.(Resp).ErrCode, mReply.(Resp).ErrMsg, "WS WriteMsg error", conn.ws.RemoteAddr().String(), uid, platform, err.Error())
	}
}

// 写到socket里面去
func (ws *WsServer) writeMsg(conn *Conn, msg []byte) error {
	conn.wsMutex.Lock()
	defer conn.wsMutex.Unlock()
	conn.ws.SetWriteDeadline(time.Now().Add(time.Duration(60) * time.Second))
	return conn.ws.WriteMessage(websocket.BinaryMessage, msg)
}
