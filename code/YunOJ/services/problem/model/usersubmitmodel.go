package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserSubmitModel = (*customUserSubmitModel)(nil)

type (
	// UserSubmitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserSubmitModel.
	UserSubmitModel interface {
		userSubmitModel
	}

	customUserSubmitModel struct {
		*defaultUserSubmitModel
	}
)

// NewUserSubmitModel returns a model for the database table.
func NewUserSubmitModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserSubmitModel {
	return &customUserSubmitModel{
		defaultUserSubmitModel: newUserSubmitModel(conn, c, opts...),
	}
}

func NewUserSubmit() *UserSubmit {
	return &UserSubmit{
		Language: 1,
	}
}
