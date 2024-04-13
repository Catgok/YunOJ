package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContestInfoModel = (*customContestInfoModel)(nil)

type (
	// ContestInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContestInfoModel.
	ContestInfoModel interface {
		contestInfoModel
	}

	customContestInfoModel struct {
		*defaultContestInfoModel
	}
)

// NewContestInfoModel returns a model for the database table.
func NewContestInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContestInfoModel {
	return &customContestInfoModel{
		defaultContestInfoModel: newContestInfoModel(conn, c, opts...),
	}
}
