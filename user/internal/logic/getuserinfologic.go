package logic

import (
	"context"

	"zerorpc_demo/user/internal/svc"
	"zerorpc_demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.IdRequest) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	i := users[int(in.Id)]
	return &user.UserResponse{Id: i.id, Name: i.name, Desc: i.desc}, nil
}
