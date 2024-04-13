package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContestProblemInfoModel = (*customContestProblemInfoModel)(nil)

type (
	// ContestProblemInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContestProblemInfoModel.
	ContestProblemInfoModel interface {
		contestProblemInfoModel
	}

	customContestProblemInfoModel struct {
		*defaultContestProblemInfoModel
	}
)

// NewContestProblemInfoModel returns a model for the database table.
func NewContestProblemInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContestProblemInfoModel {
	return &customContestProblemInfoModel{
		defaultContestProblemInfoModel: newContestProblemInfoModel(conn, c, opts...),
	}
}
