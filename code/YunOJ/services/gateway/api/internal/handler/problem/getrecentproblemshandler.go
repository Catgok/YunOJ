package problem

import (
	"net/http"

	"YunOJ/services/gateway/api/internal/logic/problem"
	"YunOJ/services/gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRecentProblemsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := problem.NewGetRecentProblemsLogic(r.Context(), svcCtx)
		resp, err := l.GetRecentProblems()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
