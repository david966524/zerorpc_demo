package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"zerorpc_demo/domain/domainclient"
	"zerorpc_demo/gateway/internal/config"
	"zerorpc_demo/user/userclient"
)

type ServiceContext struct {
	Config    config.Config
	UserRpc   userclient.User
	DomainRpc domainclient.Domain
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserRpc:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		DomainRpc: domainclient.NewDomain(zrpc.MustNewClient(c.DomainRpc)),
	}
}
