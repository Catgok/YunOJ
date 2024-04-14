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
	contestSubmitInfoFieldNames          = builder.RawFieldNames(&ContestSubmitInfo{})
	contestSubmitInfoRows                = strings.Join(contestSubmitInfoFieldNames, ",")
	contestSubmitInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(contestSubmitInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	contestSubmitInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(contestSubmitInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheContestSubmitInfoIdPrefix                = "cache:contestSubmitInfo:id:"
	cacheContestSubmitInfoContestIdSubmitIdPrefix = "cache:contestSubmitInfo:contestId:submitId:"
	cacheContestSubmitInfoSubmitIdPrefix          = "cache:contestSubmitInfo:submitId:"
)

type (
	contestSubmitInfoModel interface {
		Insert(ctx context.Context, data *ContestSubmitInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContestSubmitInfo, error)
		FindOneByContestIdSubmitId(ctx context.Context, contestId int64, submitId int64) (*ContestSubmitInfo, error)
		FindOneBySubmitId(ctx context.Context, submitId int64) (*ContestSubmitInfo, error)
		Update(ctx context.Context, data *ContestSubmitInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContestSubmitInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ContestSubmitInfo struct {
		Id        int64     `db:"id"`         // 竞赛提交信息ID
		ContestId int64     `db:"contest_id"` // 竞赛ID
		SubmitId  int64     `db:"submit_id"`  // 提交ID
		CreatedAt time.Time `db:"created_at"` // 记录创建时间
		UpdatedAt time.Time `db:"updated_at"` // 记录更新时间
	}
)

func newContestSubmitInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultContestSubmitInfoModel {
	return &defaultContestSubmitInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`contest_submit_info`",
	}
}

func (m *defaultContestSubmitInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	contestSubmitInfoContestIdSubmitIdKey := fmt.Sprintf("%s%v:%v", cacheContestSubmitInfoContestIdSubmitIdPrefix, data.ContestId, data.SubmitId)
	contestSubmitInfoIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoIdPrefix, id)
	contestSubmitInfoSubmitIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoSubmitIdPrefix, data.SubmitId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, contestSubmitInfoContestIdSubmitIdKey, contestSubmitInfoIdKey, contestSubmitInfoSubmitIdKey)
	return err
}

func (m *defaultContestSubmitInfoModel) FindOne(ctx context.Context, id int64) (*ContestSubmitInfo, error) {
	contestSubmitInfoIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoIdPrefix, id)
	var resp ContestSubmitInfo
	err := m.QueryRowCtx(ctx, &resp, contestSubmitInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contestSubmitInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultContestSubmitInfoModel) FindOneByContestIdSubmitId(ctx context.Context, contestId int64, submitId int64) (*ContestSubmitInfo, error) {
	contestSubmitInfoContestIdSubmitIdKey := fmt.Sprintf("%s%v:%v", cacheContestSubmitInfoContestIdSubmitIdPrefix, contestId, submitId)
	var resp ContestSubmitInfo
	err := m.QueryRowIndexCtx(ctx, &resp, contestSubmitInfoContestIdSubmitIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `contest_id` = ? and `submit_id` = ? limit 1", contestSubmitInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, contestId, submitId); err != nil {
			return nil, err
		}
		return resp.Id, nil
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

func (m *defaultContestSubmitInfoModel) FindOneBySubmitId(ctx context.Context, submitId int64) (*ContestSubmitInfo, error) {
	contestSubmitInfoSubmitIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoSubmitIdPrefix, submitId)
	var resp ContestSubmitInfo
	err := m.QueryRowIndexCtx(ctx, &resp, contestSubmitInfoSubmitIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `submit_id` = ? limit 1", contestSubmitInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, submitId); err != nil {
			return nil, err
		}
		return resp.Id, nil
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

func (m *defaultContestSubmitInfoModel) Insert(ctx context.Context, data *ContestSubmitInfo) (sql.Result, error) {
	contestSubmitInfoContestIdSubmitIdKey := fmt.Sprintf("%s%v:%v", cacheContestSubmitInfoContestIdSubmitIdPrefix, data.ContestId, data.SubmitId)
	contestSubmitInfoIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoIdPrefix, data.Id)
	contestSubmitInfoSubmitIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoSubmitIdPrefix, data.SubmitId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, contestSubmitInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ContestId, data.SubmitId)
	}, contestSubmitInfoContestIdSubmitIdKey, contestSubmitInfoIdKey, contestSubmitInfoSubmitIdKey)
	return ret, err
}

func (m *defaultContestSubmitInfoModel) Update(ctx context.Context, newData *ContestSubmitInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	contestSubmitInfoContestIdSubmitIdKey := fmt.Sprintf("%s%v:%v", cacheContestSubmitInfoContestIdSubmitIdPrefix, data.ContestId, data.SubmitId)
	contestSubmitInfoIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoIdPrefix, data.Id)
	contestSubmitInfoSubmitIdKey := fmt.Sprintf("%s%v", cacheContestSubmitInfoSubmitIdPrefix, data.SubmitId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, contestSubmitInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.ContestId, newData.SubmitId, newData.Id)
	}, contestSubmitInfoContestIdSubmitIdKey, contestSubmitInfoIdKey, contestSubmitInfoSubmitIdKey)
	return err
}

func (m *defaultContestSubmitInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheContestSubmitInfoIdPrefix, primary)
}

func (m *defaultContestSubmitInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contestSubmitInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultContestSubmitInfoModel) tableName() string {
	return m.table
}
