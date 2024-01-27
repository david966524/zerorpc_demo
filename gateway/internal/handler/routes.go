// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	domain "zerorpc_demo/gateway/internal/handler/domain"
	login "zerorpc_demo/gateway/internal/handler/login"
	user "zerorpc_demo/gateway/internal/handler/user"
	"zerorpc_demo/gateway/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/domain/all",
				Handler: domain.GetDomainInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: login.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/info",
				Handler: user.GetUserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/user/info/:id",
				Handler: user.DeleteUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/info/update",
				Handler: user.UpdateUserInfoHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1"),
	)
}