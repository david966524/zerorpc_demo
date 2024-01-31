package logic

import (
	"context"
	"fmt"

	"zerorpc_demo/message/internal/svc"
	"zerorpc_demo/message/rpc/types/message"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *message.SendRequest) (*message.SendResponse, error) {
	// todo: add your logic here and delete this line
	logx.Info(fmt.Sprintf("%v 短信已发送", in.Username))
	return &message.SendResponse{Msg: fmt.Sprintf("%v 短信已发送", in.Username)}, nil
}
