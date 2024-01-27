package logic

import (
	"context"
	"time"

	"zerorpc_demo/domain/internal/svc"
	"zerorpc_demo/domain/rpc/types/domain"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDomainAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDomainAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDomainAllLogic {
	return &GetDomainAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDomainAllLogic) GetDomainAll(in *domain.DomainReq) (*domain.DomainInfoResp, error) {
	// todo: add your logic here and delete this line
	logx.Info("--------------------------------domain rpc-----------")
	now := time.Now()
	zones := getLocalZones()
	logx.Info(time.Since(now))
	return &domain.DomainInfoResp{DomainList: zones}, nil
}

func getLocalZones() []*domain.DomainInfo {
	var domainList []*domain.DomainInfo
	zone1 := &domain.DomainInfo{
		Id:         "35ef8178527791a01fa4f0771020e7ad",
		DomainName: "hu85254.cc",
		Status:     "active",
	}
	zone2 := &domain.DomainInfo{
		Id:         "35ef8178527791a01fa4f0771020e7ad",
		DomainName: "hu85254.com",
		Status:     "active",
	}
	zone3 := &domain.DomainInfo{
		Id:         "35ef8178527791a01fa4f0771020e7ad",
		DomainName: "hu85254.info",
		Status:     "active",
	}
	domainList = append(domainList, zone1)
	domainList = append(domainList, zone2)
	domainList = append(domainList, zone3)
	return domainList
}
