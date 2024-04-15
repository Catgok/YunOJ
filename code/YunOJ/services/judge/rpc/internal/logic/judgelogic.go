package logic

import (
	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/internal/utils"
	"YunOJ/services/judge/rpc/judge"
	"YunOJ/services/problem/rpc/problem"
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type JudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JudgeLogic {
	return &JudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JudgeLogic) Judge(in *judge.JudgeRequest) (*judge.JudgeResponse, error) {
	resp := &judge.JudgeResponse{
		Code:    0,
		Message: "success",
	}
	// 1、调用 problem.rpc 服务 查询时空限制
	res, err := l.svcCtx.ProblemRpc.GetProblemById(l.ctx, &problem.GetProblemByIdRequest{
		ProblemId: in.ProblemId,
	})
	if err != nil {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}
	timeLimit, memLimit := res.GetProblem().GetTimeLimit(), res.GetProblem().GetMemoryLimit()

	// 2、获取input、output case
	paths, err := l.svcCtx.OssConfig.GetDirectoryFilesByProblemId(in.ProblemId)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	var inputs, outputs []string
	for _, v := range paths {
		if strings.HasSuffix(v, "in") {
			inputs = append(inputs, v)
		} else if strings.HasSuffix(v, "out") {
			outputs = append(outputs, v)
		}
	}
	if len(inputs) != len(outputs) {
		resp.Code, resp.Message = 6002, "input case number not equal to output case number"
		return resp, nil
	}
	sort.Strings(inputs)
	sort.Strings(outputs)

	inputsCase, err := l.svcCtx.OssConfig.DownloadStrings(inputs)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}
	outputsCase, err := l.svcCtx.OssConfig.DownloadStrings(outputs)
	if err != nil {
		resp.Code, resp.Message = 6001, err.Error()
		return resp, nil
	}

	// 3、for 每一个case go-judge api评测, 对比输出结果
	timeTotUsed, memTotUsed := int64(0), int64(0)
	for i := 0; i < len(inputsCase); i++ {
		res, timeUsed, memUsed := judgeOne(in.Code, inputsCase[i], outputsCase[i], in.SubmitId, timeLimit, memLimit)
		timeTotUsed = max(timeTotUsed, timeUsed)
		memTotUsed = max(memTotUsed, memUsed)
		if res != Accepted {
			//4、调用problem.rpc更新评测结果
			updateSubmitResult, err := l.svcCtx.ProblemRpc.UpdateSubmit(l.ctx, &problem.UpdateSubmitRequest{
				Submit: &problem.Submit{
					SubmitId: in.SubmitId,
					Time:     timeTotUsed,
					Memory:   memTotUsed,
					Result:   int64(res),
				},
			})
			if err != nil || updateSubmitResult.Success == false {
				resp.Code, resp.Message = 500, err.Error()
				return resp, nil
			}
			return resp, nil
		}
	}

	updateSubmitResult, err := l.svcCtx.ProblemRpc.UpdateSubmit(l.ctx, &problem.UpdateSubmitRequest{
		Submit: &problem.Submit{
			SubmitId: in.SubmitId,
			Time:     timeTotUsed,
			Memory:   memTotUsed,
			Result:   int64(Accepted),
		},
	})
	if err != nil || updateSubmitResult.Success == false {
		resp.Code, resp.Message = 500, err.Error()
		return resp, nil
	}

	return resp, nil
}

func judgeOne(code, input, output string, submitId, timeLimit, memLimit int64) (res judgeResult, timeUsed, memUsed int64) {
	timeLimitNs, memLimitB := timeLimit*1000000, memLimit<<20

	// https://github.com/criyle/go-judge/blob/master/README.cn.md

	compileReq := fmt.Sprintf(`{"cmd":[{"args":["/usr/bin/g++","%d.cpp","-o","%d"],"env":["PATH=/usr/bin:/bin"],"files":[{"content":""},{"name":"stdout","max":10240},{"name":"stderr","max":10240}],"cpuLimit":10000000000,"memoryLimit":134217728,"procLimit":50,"copyIn":{"%d.cpp":{"content":%s}},"copyOut":["stdout","stderr"],"copyOutCached":["%d"]}]}`, submitId, submitId, submitId, strconv.Quote(code), submitId)
	compileResp, err := utils.HttpPost("http://serverhost:5050/run", []byte(compileReq))
	if err != nil {
		return SystemError, 0, 0
	}
	//logx.Infof("judgeOne respBody: %s", compileResp)
	var compileData []compileResultBody
	err = json.Unmarshal(compileResp, &compileData)
	if err != nil {
		return SystemError, 0, 0
	}
	if compileData[0].Status != "Accepted" || compileData[0].ExitStatus != 0 {
		return CompileError, 0, 0
	}
	fileId := compileData[0].FileIds[fmt.Sprintf("%d", submitId)]

	runReq := fmt.Sprintf(`{"cmd":[{"args":["%d"],"env":["PATH=/usr/bin:/bin"],"files":[{"content":%s},{"name":"stdout","max":10240},{"name":"stderr","max":10240}],"cpuLimit":%d,"memoryLimit":%d,"procLimit":50,"copyIn":{"%d":{"fileId":"%s"}}}]}`, submitId, strconv.Quote(input), timeLimitNs, memLimitB, submitId, fileId)
	runResp, err := utils.HttpPost("http://serverhost:5050/run", []byte(runReq))
	if err != nil {
		return SystemError, 0, 0
	}
	logx.Infof("judgeOne respBody: %s", runResp)
	var runData []runResultBody
	err = json.Unmarshal(runResp, &runData)
	if err != nil {
		return SystemError, 0, 0
	}

	if runData[0].ExitStatus != 0 {
		return RuntimeError, 0, 0
	}
	if runData[0].RunTime > timeLimitNs {
		return TimeLimitExceeded, runData[0].RunTime, runData[0].Memory
	}
	if runData[0].Memory > memLimitB {
		return MemoryLimitExceeded, runData[0].RunTime, runData[0].Memory
	}

	reg := regexp.MustCompile(`[[:cntrl:]]`)
	cleanStdout := reg.ReplaceAllString(runData[0].Files.Stdout, "")
	cleanOutput := reg.ReplaceAllString(output, "")
	if cleanStdout != cleanOutput {
		return WrongAnswer, runData[0].RunTime, runData[0].Memory
	}
	return Accepted, runData[0].RunTime, runData[0].Memory
}

type judgeResult int

// 评测结果
const (
	Accepted            judgeResult = 0
	WrongAnswer         judgeResult = 1
	TimeLimitExceeded   judgeResult = 2
	MemoryLimitExceeded judgeResult = 3
	RuntimeError        judgeResult = 4
	OutputLimitExceeded judgeResult = 5
	CompileError        judgeResult = 6
	SystemError         judgeResult = 7
)

func (s judgeResult) GetMsg() string {
	switch s {
	case Accepted:
		return "Accepted"
	case WrongAnswer:
		return "Wrong Answer"
	case TimeLimitExceeded:
		return "Time Limit Exceeded"
	case MemoryLimitExceeded:
		return "Memory Limit Exceeded"
	case RuntimeError:
		return "Runtime Error"
	case OutputLimitExceeded:
		return "Output Limit Exceeded"
	case CompileError:
		return "Compile Error"
	case SystemError:
		return "System Error"
	default:
		return "Unknown"
	}
}

type compileResultBody struct {
	Status     string            `json:"status"`
	ExitStatus int               `json:"exitStatus"`
	FileIds    map[string]string `json:"fileIds"`
}

type runResultBody struct {
	Status     string `json:"status"`
	ExitStatus int    `json:"exitStatus"`
	Time       int64  `json:"time"`
	Memory     int64  `json:"memory"`
	RunTime    int64  `json:"runTime"`
	Files      struct {
		Stderr string `json:"stderr"`
		Stdout string `json:"stdout"`
	} `json:"files"`
}
