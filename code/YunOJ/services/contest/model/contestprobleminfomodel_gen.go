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
	contestProblemInfoFieldNames          = builder.RawFieldNames(&ContestProblemInfo{})
	contestProblemInfoRows                = strings.Join(contestProblemInfoFieldNames, ",")
	contestProblemInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(contestProblemInfoFieldNames, "`contest_problem_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	contestProblemInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(contestProblemInfoFieldNames, "`contest_problem_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheContestProblemInfoContestProblemIdPrefix   = "cache:contestProblemInfo:contestProblemId:"
	cacheContestProblemInfoContestIdProblemIdPrefix = "cache:contestProblemInfo:contestId:problemId:"
)

type (
	contestProblemInfoModel interface {
		Insert(ctx context.Context, data *ContestProblemInfo) (sql.Result, error)
		InsertBatch(ctx context.Context, data []*ContestProblemInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContestProblemInfo, error)
		FindOneByContestIdProblemId(ctx context.Context, contestId int64, problemId int64) (*ContestProblemInfo, error)
		FindByContestId(ctx context.Context, contestId int64) ([]int64, error)
		Update(ctx context.Context, data *ContestProblemInfo) error
		Delete(ctx context.Context, contestProblemId int64) error
	}

	defaultContestProblemInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ContestProblemInfo struct {
		ContestProblemId int64     `db:"contest_problem_id"` // 竞赛题目信息表ID
		ContestId        int64     `db:"contest_id"`         // 竞赛ID
		ProblemId        int64     `db:"problem_id"`         // 题目ID
		CreatedAt        time.Time `db:"created_at"`         // 记录创建时间
		UpdatedAt        time.Time `db:"updated_at"`         // 记录更新时间
	}
)

func newContestProblemInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultContestProblemInfoModel {
	return &defaultContestProblemInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`contest_problem_info`",
	}
}

func (m *defaultContestProblemInfoModel) Delete(ctx context.Context, contestProblemId int64) error {
	data, err := m.FindOne(ctx, contestProblemId)
	if err != nil {
		return err
	}

	contestProblemInfoContestIdProblemIdKey := fmt.Sprintf("%s%v:%v", cacheContestProblemInfoContestIdProblemIdPrefix, data.ContestId, data.ProblemId)
	contestProblemInfoContestProblemIdKey := fmt.Sprintf("%s%v", cacheContestProblemInfoContestProblemIdPrefix, contestProblemId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `contest_problem_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, contestProblemId)
	}, contestProblemInfoContestIdProblemIdKey, contestProblemInfoContestProblemIdKey)
	return err
}

func (m *defaultContestProblemInfoModel) FindOne(ctx context.Context, contestProblemId int64) (*ContestProblemInfo, error) {
	contestProblemInfoContestProblemIdKey := fmt.Sprintf("%s%v", cacheContestProblemInfoContestProblemIdPrefix, contestProblemId)
	var resp ContestProblemInfo
	err := m.QueryRowCtx(ctx, &resp, contestProblemInfoContestProblemIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `contest_problem_id` = ? limit 1", contestProblemInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, contestProblemId)
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

func (m *defaultContestProblemInfoModel) FindByContestId(ctx context.Context, contestId int64) ([]int64, error) {
	var resp []int64
	query := fmt.Sprintf("select problem_id from %s where `contest_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, contestId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContestProblemInfoModel) FindOneByContestIdProblemId(ctx context.Context, contestId int64, problemId int64) (*ContestProblemInfo, error) {
	contestProblemInfoContestIdProblemIdKey := fmt.Sprintf("%s%v:%v", cacheContestProblemInfoContestIdProblemIdPrefix, contestId, problemId)
	var resp ContestProblemInfo
	err := m.QueryRowIndexCtx(ctx, &resp, contestProblemInfoContestIdProblemIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `contest_id` = ? and `problem_id` = ? limit 1", contestProblemInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, contestId, problemId); err != nil {
			return nil, err
		}
		return resp.ContestProblemId, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultContestProblemInfoModel) Insert(ctx context.Context, data *ContestProblemInfo) (sql.Result, error) {
	contestProblemInfoContestIdProblemIdKey := fmt.Sprintf("%s%v:%v", cacheContestProblemInfoContestIdProblemIdPrefix, data.ContestId, data.ProblemId)
	contestProblemInfoContestProblemIdKey := fmt.Sprintf("%s%v", cacheContestProblemInfoContestProblemIdPrefix, data.ContestProblemId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, contestProblemInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ContestId, data.ProblemId)
	}, contestProblemInfoContestIdProblemIdKey, contestProblemInfoContestProblemIdKey)
	return ret, err
}

func (m *defaultContestProblemInfoModel) InsertBatch(ctx context.Context, data []*ContestProblemInfo) (sql.Result, error) {
	if len(data) == 0 {
		return nil, nil
	}
	query := fmt.Sprintf("insert into %s (%s) values ", m.table, contestProblemInfoRowsExpectAutoSet)
	vals := []interface{}{}
	for _, item := range data {
		vals = append(vals, item.ContestId, item.ProblemId)
		query += "(?, ?),"
	}
	query = query[0 : len(query)-1]
	return m.ExecNoCache(query, vals...)
}

func (m *defaultContestProblemInfoModel) Update(ctx context.Context, newData *ContestProblemInfo) error {
	data, err := m.FindOne(ctx, newData.ContestProblemId)
	if err != nil {
		return err
	}

	contestProblemInfoContestIdProblemIdKey := fmt.Sprintf("%s%v:%v", cacheContestProblemInfoContestIdProblemIdPrefix, data.ContestId, data.ProblemId)
	contestProblemInfoContestProblemIdKey := fmt.Sprintf("%s%v", cacheContestProblemInfoContestProblemIdPrefix, data.ContestProblemId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `contest_problem_id` = ?", m.table, contestProblemInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ContestId, newData.ProblemId, newData.ContestProblemId)
	}, contestProblemInfoContestIdProblemIdKey, contestProblemInfoContestProblemIdKey)
	return err
}

func (m *defaultContestProblemInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheContestProblemInfoContestProblemIdPrefix, primary)
}

func (m *defaultContestProblemInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `contest_problem_id` = ? limit 1", contestProblemInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultContestProblemInfoModel) tableName() string {
	return m.table
}
