package mq

import (
	"YunOJ/services/contest/rpc/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitChange struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitChange(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitChange {
	return &SubmitChange{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitChange) Consume(key, val string) error {
	logx.Infof("Consume key :%s , val :%s", key, val)
	// todo 消费逻辑
	return nil
}
