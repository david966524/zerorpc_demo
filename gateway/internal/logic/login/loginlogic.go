package login

import (
	"context"
	"zerorpc_demo/user/rpc/types/user"

	"zerorpc_demo/gateway/internal/svc"
	"zerorpc_demo/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	userInfoResp, err := l.svcCtx.UserRpc.Login(l.ctx, &user.UserLogin{
		Username: req.Username,
		Password: req.Password,
	})
	logx.Info(req.Enduser)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		Id:    userInfoResp.GetId(),
		Name:  userInfoResp.Name,
		Token: "xxxx.xxxx.xxxx.xxxx",
	}, nil
}
