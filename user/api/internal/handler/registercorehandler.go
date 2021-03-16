package handler

import (
	"net/http"

	"yl/user/api/internal/logic"
	"yl/user/api/internal/svc"
	"yl/user/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func registerCoreHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterCoreReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		l := logic.NewRegisterCoreLogic(r.Context(), ctx)
		resp, err := l.RegisterCore(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
