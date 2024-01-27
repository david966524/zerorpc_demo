package logic

import (
	"context"
	"errors"
	"fmt"
	"zerorpc_demo/email/rpc/types/email"

	"zerorpc_demo/user/internal/svc"
	"zerorpc_demo/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.UserLogin) (*user.UserResponse, error) {
	// todo: add your logic here and delete this line
	if in.Username != "david" || in.Password != "123456" {
		return &user.UserResponse{}, errors.New("不对")
	}
	sendEmail, _ := l.svcCtx.EmailRpc.SendEmail(l.ctx, &email.UserEmailRequest{
		Username: in.Username,
		Email:    fmt.Sprintf("%v@gmail.com", in.Username),
	})
	logx.Info(sendEmail.Msg)
	return &user.UserResponse{
		Id:   100,
		Name: in.Username,
		Desc: "rpc return",
	}, nil
}
