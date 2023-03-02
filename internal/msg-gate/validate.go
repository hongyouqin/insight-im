package msggate

import (
	"insight/pkg/common/constant"
	"insight/pkg/protos/msg"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
)

//websocket 上行，下行结构

type Req struct {
	ReqIdentifier int32  `json:"reqIdentifier" validate:"required"`
	Token         string `json:"token" `
	SendID        string `json:"sendID" validate:"required"`
	OperationID   string `json:"operationID" validate:"required"`
	MsgIncr       string `json:"msgIncr" validate:"required"`
	Data          []byte `json:"data"`
}

type Resp struct {
	ReqIdentifier int32  `json:"reqIdentifier"`
	MsgIncr       string `json:"msgIncr"`
	OperationID   string `json:"operationID"`
	ErrCode       int32  `json:"errCode"`
	ErrMsg        string `json:"errMsg"`
	Data          []byte `json:"data"`
}

// 上行数据验证
func (ws *WsServer) argsValidate(req *Req, indetifier int32) (isPass bool, errCode int32, errMsg string, returnData interface{}) {
	switch indetifier {
	case constant.WSSendMsg:
		data := msg.MsgData{}
		if err := proto.Unmarshal(req.Data, &data); err != nil {
			ws.log.Error("unmarshal data struct err", zap.String("errr", err.Error()), zap.Int32("indetifier", indetifier))
			return false, 203, err.Error(), nil
		}
		if err := ws.validate.Struct(data); err != nil {
			ws.log.Error("data args validate err", zap.String("errr", err.Error()), zap.Int32("indetifier", indetifier))
			return false, 204, err.Error(), nil

		}
		return true, 0, "", data
	}
	return false, 204, "input args err", nil
}
