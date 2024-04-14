package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContestSubmitInfoModel = (*customContestSubmitInfoModel)(nil)

type (
	// ContestSubmitInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContestSubmitInfoModel.
	ContestSubmitInfoModel interface {
		contestSubmitInfoModel
	}

	customContestSubmitInfoModel struct {
		*defaultContestSubmitInfoModel
	}
)

// NewContestSubmitInfoModel returns a model for the database table.
func NewContestSubmitInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContestSubmitInfoModel {
	return &customContestSubmitInfoModel{
		defaultContestSubmitInfoModel: newContestSubmitInfoModel(conn, c, opts...),
	}
}
