// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"
	"time"

	contest "YunOJ/services/gateway/api/internal/handler/contest"
	judge "YunOJ/services/gateway/api/internal/handler/judge"
	problem "YunOJ/services/gateway/api/internal/handler/problem"
	user "YunOJ/services/gateway/api/internal/handler/user"
	"YunOJ/services/gateway/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/getContestById",
				Handler: contest.GetContestByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getContestProblems",
				Handler: contest.GetContestProblemsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getContestRank",
				Handler: contest.GetContestRankHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getContestsByPage",
				Handler: contest.GetContestsByPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getRecentContests",
				Handler: contest.GetRecentContestsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/contest"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/addProblem",
					Handler: contest.AddProblemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/createContest",
					Handler: contest.CreateContestHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/getSignUpContests",
					Handler: contest.GetSignUpContestsHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/signUpContest",
					Handler: contest.SignUpContestHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/submitAnswer",
					Handler: contest.SubmitAnswerHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/updateContest",
					Handler: contest.UpdateContestHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/contest"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/onlineJudge",
				Handler: judge.OnlineJudgeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/judge"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/setJudgeCase",
					Handler: judge.SetJudgeCaseHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/judge"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/get",
				Handler: problem.GetProblemByIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getByPage",
				Handler: problem.GetProblemsByPageHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/getRecentProblems",
				Handler: problem.GetRecentProblemsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/problem"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/create",
					Handler: problem.CreateProblemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/delete",
					Handler: problem.DeleteProblemHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/getSubmissionByProblemId",
					Handler: problem.GetSubmissionByUserIdAndProblemIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/getSubmit",
					Handler: problem.GetSubmitByIdHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/submit",
					Handler: problem.SubmitHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: problem.UpdateProblemHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/problem"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/v1/user"),
		rest.WithTimeout(3000*time.Millisecond),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.AuthInterceptor},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/getUserInfoByUToken",
					Handler: user.GetUserInfoByUTokenHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/user"),
		rest.WithTimeout(3000*time.Millisecond),
	)
}
