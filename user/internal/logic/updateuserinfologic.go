package logic

import (
	"context"

	"zerorpc_demo/user/internal/svc"
	"zerorpc_demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *user.UserInfo) (*user.ResponseMsg, error) {
	// todo: add your logic here and delete this line

	return &user.ResponseMsg{}, nil
}
