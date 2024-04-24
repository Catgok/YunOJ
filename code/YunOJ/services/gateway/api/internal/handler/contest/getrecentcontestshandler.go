package contest

import (
	"net/http"

	"YunOJ/services/gateway/api/internal/logic/contest"
	"YunOJ/services/gateway/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRecentContestsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := contest.NewGetRecentContestsLogic(r.Context(), svcCtx)
		resp, err := l.GetRecentContests()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
