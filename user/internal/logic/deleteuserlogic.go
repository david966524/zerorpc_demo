package logic

import (
	"context"
	"fmt"
	"zerorpc_demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"zerorpc_demo/user/internal/svc"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *user.IdRequest) (*user.ResponseMsg, error) {
	// todo: add your logic here and delete this line

	delete(users, int(in.Id))
	return &user.ResponseMsg{
		Code: 200,
		Msg:  fmt.Sprintf("user id = %v, delete", in.Id),
	}, nil
}
