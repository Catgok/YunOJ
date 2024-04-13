package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ContestUserInfoModel = (*customContestUserInfoModel)(nil)

type (
	// ContestUserInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customContestUserInfoModel.
	ContestUserInfoModel interface {
		contestUserInfoModel
	}

	customContestUserInfoModel struct {
		*defaultContestUserInfoModel
	}
)

// NewContestUserInfoModel returns a model for the database table.
func NewContestUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ContestUserInfoModel {
	return &customContestUserInfoModel{
		defaultContestUserInfoModel: newContestUserInfoModel(conn, c, opts...),
	}
}
