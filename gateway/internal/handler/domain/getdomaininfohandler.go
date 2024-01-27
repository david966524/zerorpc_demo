package domain

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"zerorpc_demo/gateway/internal/logic/domain"
	"zerorpc_demo/gateway/internal/svc"
)

func GetDomainInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := domain.NewGetDomainInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetDomainInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
