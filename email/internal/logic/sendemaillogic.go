package logic

import (
	"context"
	"errors"
	"fmt"

	"zerorpc_demo/email/internal/svc"
	"zerorpc_demo/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendEmailLogic {
	return &SendEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendEmailLogic) SendEmail(in *email.UserEmailRequest) (*email.UserEmailResponse, error) {
	// todo: add your logic here and delete this line
	var err error = nil
	if in.Email == "" {
		err = errors.New("email 不能为空")
	}
	return &email.UserEmailResponse{Msg: fmt.Sprintf("向用户 %v ，邮箱为 %v ，发送验证邮件", in.Username, in.Email)}, err
}
