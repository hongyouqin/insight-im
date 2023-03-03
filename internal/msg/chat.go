package msg

import (
	"context"
	"insight/internal/kafka"
	"insight/pkg/common/constant"
	"insight/pkg/proto/msg"
	rpc "insight/pkg/proto/msg"

	"go.uber.org/zap"
)

// 投递消息到mq
// 用户关系验证
type Chat struct {
	producer *kafka.Producer
	log      *zap.Logger
	rpc.UnimplementedChatServer
}

func NewChatServer(log *zap.Logger) *Chat {
	chat := Chat{
		producer: kafka.NewKafkaProducer("127.0.0.1:9092", "ws2ms_chat"),
		log:      log,
	}
	return &chat
}

func (c *Chat) SendMsg(ctx context.Context, req *msg.SendMsgReq) (*msg.SendMsgResp, error) {

	//token 验证
	resp := msg.SendMsgResp{}

	switch req.Data.SessionType {
	case constant.SingleChatType:
		//接收者mq
		err := c.deliverMsgToKafka(req, req.Data.RecvID)
		if err != nil {
			c.log.Error("kfka send msg err", zap.String("recvId", req.Data.RecvID), zap.String("msg", req.String()))
			return returnMsg(&resp, req, 201, "kfka send msg err", "", 0)
		}
		//发送者存mq, 排除自己
		if req.Data.SendID != req.Data.RecvID {
			err = c.deliverMsgToKafka(req, req.Data.SendID)
			c.log.Error("kfka send msg err", zap.String("sendId", req.Data.SendID), zap.String("msg", req.String()))
			return returnMsg(&resp, req, 201, "kfka send msg err", "", 0)
		}
		return returnMsg(&resp, req, 0, "", req.Data.ServerMsgID, req.Data.SendTime)
	case constant.GroupChatType:
		//获取群成员

		//遍历群里成员
		//消息存入kafka收件箱，每一个用户都有一个自己的收件箱，收件箱使用userId来区分

	default:
		//
	}
	return returnMsg(&resp, req, 203, "unkonwn sessionType", "", 0)
}

func (c *Chat) GetMaxAndMinSeq(context.Context, *msg.GetMaxAndMinSeqReq) (*msg.GetMaxAndMinSeqResp, error) {
	return nil, nil
}

// 投递消息给kafka
func (c *Chat) deliverMsgToKafka(msg *msg.SendMsgReq, key string) error {
	pid, offset, err := c.producer.SendMessage(msg, key)
	if err != nil {
		c.log.Error("kafka send failed", zap.String("send data", msg.String()), zap.Int32("pid", pid), zap.Int64("offset", offset), zap.String("err", err.Error()), zap.String("key", key))
		return err
	}
	return nil
}

func returnMsg(replay *msg.SendMsgResp, req *msg.SendMsgReq, errCode int32, errMsg, serverMsgID string, sendTime int64) (*msg.SendMsgResp, error) {
	replay.ErrCode = errCode
	replay.ErrMsg = errMsg
	replay.ServerMsgID = serverMsgID
	replay.ClientMsgID = req.Data.ClientMsgID
	replay.SendTime = sendTime
	return replay, nil
}
