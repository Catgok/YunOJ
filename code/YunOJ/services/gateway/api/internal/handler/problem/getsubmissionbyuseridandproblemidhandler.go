package problem

import (
	"net/http"

	"YunOJ/services/gateway/api/internal/logic/problem"
	"YunOJ/services/gateway/api/internal/svc"
	"YunOJ/services/gateway/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSubmissionByUserIdAndProblemIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetSubmissionByProblemIdRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := problem.NewGetSubmissionByUserIdAndProblemIdLogic(r.Context(), svcCtx)
		resp, err := l.GetSubmissionByUserIdAndProblemId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
