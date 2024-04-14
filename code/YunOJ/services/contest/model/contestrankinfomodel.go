package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContestRankInfoModel = (*customContestRankInfoModel)(nil)

type (
	// ContestRankInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContestRankInfoModel.
	ContestRankInfoModel interface {
		contestRankInfoModel
	}

	customContestRankInfoModel struct {
		*defaultContestRankInfoModel
	}
)

// NewContestRankInfoModel returns a model for the database table.
func NewContestRankInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContestRankInfoModel {
	return &customContestRankInfoModel{
		defaultContestRankInfoModel: newContestRankInfoModel(conn, c, opts...),
	}
}
