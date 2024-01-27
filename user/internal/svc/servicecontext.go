package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerorpc_demo/email/emailclient"
	"zerorpc_demo/email/rpc/types/email"
	"zerorpc_demo/user/internal/config"
)

type ServiceContext struct {
	Config   config.Config
	EmailRpc email.EmailClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		EmailRpc: emailclient.NewEmail(zrpc.MustNewClient(c.EmailRpc)),
	}
}
