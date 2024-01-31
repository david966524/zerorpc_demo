// Code generated by goctl. DO NOT EDIT.
// Source: message.proto

package messageclient

import (
	"context"

	"zerorpc_demo/message/rpc/types/message"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	SendRequest  = message.SendRequest
	SendResponse = message.SendResponse

	Message interface {
		SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	}

	defaultMessage struct {
		cli zrpc.Client
	}
)

func NewMessage(cli zrpc.Client) Message {
	return &defaultMessage{
		cli: cli,
	}
}

func (m *defaultMessage) SendMessage(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	client := message.NewMessageClient(m.cli.Conn())
	return client.SendMessage(ctx, in, opts...)
}
