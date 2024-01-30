package user

import (
	"context"
	"zerorpc_demo/user/rpc/types/user"

	"zerorpc_demo/gateway/internal/svc"
	"zerorpc_demo/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserLogic) DeleteUser(req *types.DeleteUserInfoId) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	deleteUser, err := l.svcCtx.UserRpc.DeleteUser(l.ctx, &user.IdRequest{Id: req.Id})
	return &types.GetUserInfoResp{
		Id:   deleteUser.Code,
		Name: "",
		Desc: deleteUser.Msg,
	}, nil
}
