// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userSubmitFieldNames          = builder.RawFieldNames(&UserSubmit{})
	userSubmitRows                = strings.Join(userSubmitFieldNames, ",")
	userSubmitRowsExpectAutoSet   = strings.Join(stringx.Remove(userSubmitFieldNames, "`submit_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userSubmitRowsWithPlaceHolder = strings.Join(stringx.Remove(userSubmitFieldNames, "`submit_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserSubmitSubmitIdPrefix           = "cache:userSubmit:submitId:"
	cacheUserSubmitUserIdAndProblemIdPrefix = "cache:userSubmit:userIdAndProblemId:"
)

type (
	userSubmitModel interface {
		Insert(ctx context.Context, data *UserSubmit) (sql.Result, error)
		FindOne(ctx context.Context, submitId int64) (*UserSubmit, error)
		FindByUserIdAndProblemId(ctx context.Context, userId, problemId int64) ([]UserSubmit, error)
		Update(ctx context.Context, data *UserSubmit) error
		Delete(ctx context.Context, submitId int64) error
	}

	defaultUserSubmitModel struct {
		sqlc.CachedConn
		table string
	}

	UserSubmit struct {
		SubmitId   int64     `db:"submit_id"`   // 提交id
		UserId     int64     `db:"user_id"`     // 用户id
		ProblemId  int64     `db:"problem_id"`  // 题目id
		Code       string    `db:"code"`        // 代码
		Status     int64     `db:"status"`      // 状态 0-未评测 1-编译中 2-编译错误 3-评测中 4-评测完成
		Language   int64     `db:"language"`    // 语言 1-c++
		Result     int64     `db:"result"`      // 结果
		Time       int64     `db:"time"`        // 运行时间 单位ms
		Memory     int64     `db:"memory"`      // 内存占用 单位MB
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newUserSubmitModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserSubmitModel {
	return &defaultUserSubmitModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user_submit`",
	}
}

func (m *defaultUserSubmitModel) Delete(ctx context.Context, submitId int64) error {
	userSubmitSubmitIdKey := fmt.Sprintf("%s%v", cacheUserSubmitSubmitIdPrefix, submitId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `submit_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, submitId)
	}, userSubmitSubmitIdKey)
	return err
}

func (m *defaultUserSubmitModel) FindByUserIdAndProblemId(ctx context.Context, userId, problemId int64) ([]UserSubmit, error) {
	userSubmitUserIdAndProblemIdKey := fmt.Sprintf("%s%v%v", cacheUserSubmitUserIdAndProblemIdPrefix, userId, problemId)
	var resp []UserSubmit
	err := m.QueryRowCtx(ctx, &resp, userSubmitUserIdAndProblemIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? and `problem_id` = ? order by `update_time` ", userSubmitRows, m.table)
		return conn.QueryRowsCtx(ctx, v, query, userId, problemId)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserSubmitModel) FindOne(ctx context.Context, submitId int64) (*UserSubmit, error) {
	userSubmitSubmitIdKey := fmt.Sprintf("%s%v", cacheUserSubmitSubmitIdPrefix, submitId)
	var resp UserSubmit
	err := m.QueryRowCtx(ctx, &resp, userSubmitSubmitIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `submit_id` = ? limit 1", userSubmitRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, submitId)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserSubmitModel) Insert(ctx context.Context, data *UserSubmit) (sql.Result, error) {
	userSubmitSubmitIdKey := fmt.Sprintf("%s%v", cacheUserSubmitSubmitIdPrefix, data.SubmitId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userSubmitRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ProblemId, data.Code, data.Status, data.Language, data.Result, data.Time, data.Memory)
	}, userSubmitSubmitIdKey)
	return ret, err
}

func (m *defaultUserSubmitModel) Update(ctx context.Context, data *UserSubmit) error {
	userSubmitSubmitIdKey := fmt.Sprintf("%s%v", cacheUserSubmitSubmitIdPrefix, data.SubmitId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `submit_id` = ?", m.table, userSubmitRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ProblemId, data.Code, data.Status, data.Language, data.Result, data.Time, data.Memory, data.SubmitId)
	}, userSubmitSubmitIdKey)
	return err
}

func (m *defaultUserSubmitModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserSubmitSubmitIdPrefix, primary)
}

func (m *defaultUserSubmitModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `submit_id` = ? limit 1", userSubmitRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserSubmitModel) tableName() string {
	return m.table
}
