// Code generated by goctl. DO NOT EDIT.
package types

type BaseResponse struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type CreateProblemRequest struct {
	Title       string `json:"title"`
	TimeLimit   int64  `json:"timeLimit"`
	MemoryLimit int64  `json:"memoryLimit"`
	Description string `json:"description"`
	HardLevel   int64  `json:"hardLevel"`
}

type CreateProblemResponse struct {
	BaseResponse
	Data int64 `json:"data"`
}

type DeleteProblemRequest struct {
	ProblemId int64 `json:"problemId"`
}

type DeleteProblemResponse struct {
	BaseResponse
	Data bool `json:"data"`
}

type GetProblemByIdRequest struct {
	ProblemId int64 `json:"problemId"`
}

type GetProblemByIdResponse struct {
	BaseResponse
	Data Problem `json:"data"`
}

type GetProblemsByPageRequest struct {
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type GetProblemsByPageResponse struct {
	BaseResponse
	Data  []Problem `json:"data"`
	Total int64     `json:"total"`
}

type GetSubmissionByProblemIdRequest struct {
	UserId    int64 `json:"userId"`
	ProblemId int64 `json:"problemId"`
}

type GetSubmissionByProblemIdResponse struct {
	BaseResponse
	Data []Submit `json:"data"`
}

type GetSubmitByIdRequest struct {
	SubmitId int64 `json:"submitId"`
}

type GetSubmitByIdResponse struct {
	BaseResponse
	Data Submit `json:"data"`
}

type JudgeCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}

type LoginRequest struct {
	UserKey  string `json:"userKey"`
	Password string `json:"password"`
}

type LoginResponse struct {
	BaseResponse
	Data   User   `json:"data"`
	Utoken string `json:"utoken"`
}

type OnlineJudgeRequest struct {
	Code     string `json:"code"`
	Language int64  `json:"language"`
	Input    string `json:"input"`
}

type OnlineJudgeResponse struct {
	BaseResponse
	Data string `json:"data"`
}

type Problem struct {
	ProblemId   int64  `json:"problemId"`
	Title       string `json:"title"`
	TimeLimit   int64  `json:"timeLimit"`
	MemoryLimit int64  `json:"memoryLimit"`
	Description string `json:"description"`
	HardLevel   int64  `json:"hardLevel"`
	SubmitCount int64  `json:"submitCount"`
	PassCount   int64  `json:"passCount"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	BaseResponse
	Data bool `json:"data"`
}

type SetJudgeCaseRequest struct {
	ProblemId int64       `json:"problemId"`
	Cases     []JudgeCase `json:"cases"`
}

type SetJudgeCaseResponse struct {
	BaseResponse
}

type Submit struct {
	SubmitId  int64  `json:"submitId"`
	UserId    int64  `json:"userId"`
	ProblemId int64  `json:"problemId"`
	Code      string `json:"code"`
	Status    int64  `json:"status"`
	Language  int64  `json:"language"`
	Result    int64  `json:"result"`
	Time      int64  `json:"time"`
	Memory    int64  `json:"memory"`
}

type SubmitRequest struct {
	ProblemId int64  `json:"problemId"`
	UserId    int64  `json:"userId"`
	Code      string `json:"code"`
	Language  int64  `json:"language"`
}

type SubmitResponse struct {
	BaseResponse
	Data int64 `json:"data"`
}

type UpdateProblemRequest struct {
	ProblemId   int64  `json:"problemId"`
	Title       string `json:"title"`
	TimeLimit   int64  `json:"timeLimit"`
	MemoryLimit int64  `json:"memoryLimit"`
	Description string `json:"description"`
	HardLevel   int64  `json:"hardLevel"`
}

type UpdateProblemResponse struct {
	BaseResponse
	Data bool `json:"data"`
}

type User struct {
	UserId   int64  `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	UserType int64  `json:"userType"`
	Avatar   string `json:"avatar"`
	Status   int64  `json:"status"`
}
