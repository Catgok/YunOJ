package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ProblemModel = (*customProblemModel)(nil)

type (
	// ProblemModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProblemModel.
	ProblemModel interface {
		problemModel
	}

	customProblemModel struct {
		*defaultProblemModel
	}
)

// NewProblemModel returns a model for the database table.
func NewProblemModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ProblemModel {
	return &customProblemModel{
		defaultProblemModel: newProblemModel(conn, c, opts...),
	}
}

func NewProblem() *Problem {
	return &Problem{}
}
