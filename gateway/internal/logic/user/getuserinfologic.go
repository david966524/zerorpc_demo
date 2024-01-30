package user

import (
	"context"
	"zerorpc_demo/user/rpc/types/user"

	"zerorpc_demo/gateway/internal/svc"
	"zerorpc_demo/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	info, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.IdRequest{Id: req.Id})
	return &types.GetUserInfoResp{
		Id:   info.GetId(),
		Name: info.Name,
		Desc: info.Desc,
	}, nil
}
