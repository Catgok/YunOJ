package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JudgeCaseModel = (*customJudgeCaseModel)(nil)

type (
	// JudgeCaseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJudgeCaseModel.
	JudgeCaseModel interface {
		judgeCaseModel
	}

	customJudgeCaseModel struct {
		*defaultJudgeCaseModel
	}
)

// NewJudgeCaseModel returns a model for the database table.
func NewJudgeCaseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) JudgeCaseModel {
	return &customJudgeCaseModel{
		defaultJudgeCaseModel: newJudgeCaseModel(conn, c, opts...),
	}
}
