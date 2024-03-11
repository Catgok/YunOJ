package handler

import (
	"net/http"

	"YunOJ/services/test/api/internal/logic"
	"YunOJ/services/test/api/internal/svc"
	"YunOJ/services/test/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getOneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetOneLogic(r.Context(), svcCtx)
		resp, err := l.GetOne(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
