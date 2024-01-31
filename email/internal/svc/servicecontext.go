package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerorpc_demo/email/internal/config"
	"zerorpc_demo/message/messageclient"
	"zerorpc_demo/message/rpc/types/message"
)

type ServiceContext struct {
	Config     config.Config
	MessageRpc message.MessageClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		MessageRpc: messageclient.NewMessage(zrpc.MustNewClient(c.MessageRpc)),
	}
}
