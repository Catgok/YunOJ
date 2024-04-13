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
	contestUserInfoFieldNames          = builder.RawFieldNames(&ContestUserInfo{})
	contestUserInfoRows                = strings.Join(contestUserInfoFieldNames, ",")
	contestUserInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(contestUserInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	contestUserInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(contestUserInfoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheContestUserInfoIdPrefix              = "cache:contestUserInfo:id:"
	cacheContestUserInfoContestIdUserIdPrefix = "cache:contestUserInfo:contestId:userId:"
)

type (
	contestUserInfoModel interface {
		Insert(ctx context.Context, data *ContestUserInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ContestUserInfo, error)
		FindOneByContestIdUserId(ctx context.Context, contestId int64, userId int64) (*ContestUserInfo, error)
		Update(ctx context.Context, data *ContestUserInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultContestUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	ContestUserInfo struct {
		Id                 int64     `db:"id"`                   // 竞赛用户信息ID
		UserId             int64     `db:"user_id"`              // 用户ID
		ContestId          int64     `db:"contest_id"`           // 竞赛ID
		JoinStatus         int64     `db:"join_status"`          // 用户参加竞赛的状态 0:未参加 1:已参加 2:已结束
		ContestRank        int64     `db:"contest_rank"`         // 用户在竞赛中的排名
		ContestPassCount   int64     `db:"contest_pass_count"`   // 用户在竞赛中通过的题目数
		ContestPenaltyTime int64     `db:"contest_penalty_time"` // 用户在竞赛中的罚时
		CreatedAt          time.Time `db:"created_at"`           // 记录创建时间
		UpdatedAt          time.Time `db:"updated_at"`           // 记录更新时间
	}
)

func newContestUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultContestUserInfoModel {
	return &defaultContestUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`contest_user_info`",
	}
}

func (m *defaultContestUserInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	contestUserInfoContestIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheContestUserInfoContestIdUserIdPrefix, data.ContestId, data.UserId)
	contestUserInfoIdKey := fmt.Sprintf("%s%v", cacheContestUserInfoIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, contestUserInfoContestIdUserIdKey, contestUserInfoIdKey)
	return err
}

func (m *defaultContestUserInfoModel) FindOne(ctx context.Context, id int64) (*ContestUserInfo, error) {
	contestUserInfoIdKey := fmt.Sprintf("%s%v", cacheContestUserInfoIdPrefix, id)
	var resp ContestUserInfo
	err := m.QueryRowCtx(ctx, &resp, contestUserInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contestUserInfoRows, m.table)
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

func (m *defaultContestUserInfoModel) FindOneByContestIdUserId(ctx context.Context, contestId int64, userId int64) (*ContestUserInfo, error) {
	contestUserInfoContestIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheContestUserInfoContestIdUserIdPrefix, contestId, userId)
	var resp ContestUserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, contestUserInfoContestIdUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `contest_id` = ? and `user_id` = ? limit 1", contestUserInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, contestId, userId); err != nil {
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

func (m *defaultContestUserInfoModel) Insert(ctx context.Context, data *ContestUserInfo) (sql.Result, error) {
	contestUserInfoContestIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheContestUserInfoContestIdUserIdPrefix, data.ContestId, data.UserId)
	contestUserInfoIdKey := fmt.Sprintf("%s%v", cacheContestUserInfoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, contestUserInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ContestId, data.JoinStatus, data.ContestRank, data.ContestPassCount, data.ContestPenaltyTime)
	}, contestUserInfoContestIdUserIdKey, contestUserInfoIdKey)
	return ret, err
}

func (m *defaultContestUserInfoModel) Update(ctx context.Context, newData *ContestUserInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	contestUserInfoContestIdUserIdKey := fmt.Sprintf("%s%v:%v", cacheContestUserInfoContestIdUserIdPrefix, data.ContestId, data.UserId)
	contestUserInfoIdKey := fmt.Sprintf("%s%v", cacheContestUserInfoIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, contestUserInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.UserId, newData.ContestId, newData.JoinStatus, newData.ContestRank, newData.ContestPassCount, newData.ContestPenaltyTime, newData.Id)
	}, contestUserInfoContestIdUserIdKey, contestUserInfoIdKey)
	return err
}

func (m *defaultContestUserInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheContestUserInfoIdPrefix, primary)
}

func (m *defaultContestUserInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", contestUserInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultContestUserInfoModel) tableName() string {
	return m.table
}
