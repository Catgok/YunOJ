package logic

import (
	"YunOJ/services/judge/rpc/internal/svc"
	"YunOJ/services/judge/rpc/internal/utils"
	"YunOJ/services/judge/rpc/judge"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnlineJudgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOnlineJudgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnlineJudgeLogic {
	return &OnlineJudgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OnlineJudgeLogic) OnlineJudge(in *judge.OnlineJudgeRequest) (*judge.OnlineJudgeResponse, error) {
	resp := &judge.OnlineJudgeResponse{
		Code:    0,
		Message: "success",
	}
	randomNumber := rand.Intn(10000)
	submitId := time.Now().UnixNano()/100 + int64(randomNumber)
	var timeLimit, memLimit int64 = 1000, 512
	res, output := onlineJudgeOne(in.Code, in.Input, submitId, timeLimit, memLimit)
	if len(output) > 4096 {
		output = output[:4096]
	}
	resp.Code, resp.Message = int64(res), res.GetMsg()
	resp.Output = output
	return resp, nil
}

func onlineJudgeOne(code, input string, submitId, timeLimit, memLimit int64) (res judgeResult, output string) {
	timeLimitNs, memLimitB := timeLimit*1000000, memLimit<<20

	compileReq := fmt.Sprintf(`{"cmd":[{"args":["/usr/bin/g++","%d.cpp","-o","%d"],"env":["PATH=/usr/bin:/bin"],"files":[{"content":""},{"name":"stdout","max":10240},{"name":"stderr","max":10240}],"cpuLimit":10000000000,"memoryLimit":134217728,"procLimit":50,"copyIn":{"%d.cpp":{"content":%s}},"copyOut":["stdout","stderr"],"copyOutCached":["%d"]}]}`, submitId, submitId, submitId, strconv.Quote(code), submitId)
	compileResp, err := utils.HttpPost("http://serverhost:5050/run", []byte(compileReq))
	if err != nil {
		return SystemError, ""
	}
	//logx.Infof("judgeOne respBody: %s", compileResp)
	var compileData []compileResultBody
	err = json.Unmarshal(compileResp, &compileData)
	if err != nil {
		return SystemError, ""
	}
	if compileData[0].Status != "Accepted" || compileData[0].ExitStatus != 0 {
		return CompileError, ""
	}
	fileId := compileData[0].FileIds[fmt.Sprintf("%d", submitId)]

	runReq := fmt.Sprintf(`{"cmd":[{"args":["%d"],"env":["PATH=/usr/bin:/bin"],"files":[{"content":%s},{"name":"stdout","max":10240},{"name":"stderr","max":10240}],"cpuLimit":%d,"memoryLimit":%d,"procLimit":50,"copyIn":{"%d":{"fileId":"%s"}}}]}`, submitId, strconv.Quote(input), timeLimitNs, memLimitB, submitId, fileId)
	runResp, err := utils.HttpPost("http://serverhost:5050/run", []byte(runReq))
	if err != nil {
		return SystemError, ""
	}
	logx.Infof("judgeOne respBody: %s", runResp)
	var runData []runResultBody
	err = json.Unmarshal(runResp, &runData)
	if err != nil {
		return SystemError, ""
	}
	if runData[0].ExitStatus != 0 {
		return RuntimeError, ""
	}
	output = runData[0].Files.Stdout
	if runData[0].RunTime > timeLimitNs {
		return TimeLimitExceeded, output
	}
	if runData[0].Memory > memLimitB {
		return MemoryLimitExceeded, output
	}
	return Accepted, output
}
