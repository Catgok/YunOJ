package judge

import (
	"net/http"

	"YunOJ/services/gateway/api/internal/logic/judge"
	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func SetJudgeCaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SetJudgeCaseRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := judge.NewSetJudgeCaseLogic(r.Context(), svcCtx)
		resp, err := l.SetJudgeCase(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
