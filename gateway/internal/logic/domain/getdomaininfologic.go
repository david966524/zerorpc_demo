package domain

import (
	"context"
	"encoding/json"
	"zerorpc_demo/domain/rpc/types/domain"

	"zerorpc_demo/gateway/internal/svc"
	"zerorpc_demo/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDomainInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDomainInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDomainInfoLogic {
	return &GetDomainInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDomainInfoLogic) GetDomainInfo() (resp []types.RespDomainInfo, err error) {
	// todo: add your logic here and delete this line
	allDomain, err := l.svcCtx.DomainRpc.GetDomainAll(l.ctx, &domain.DomainReq{})
	logx.Info(allDomain.GetDomainList())
	var list []types.RespDomainInfo
	marshal, _ := json.Marshal(allDomain.GetDomainList())
	err = json.Unmarshal(marshal, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}
