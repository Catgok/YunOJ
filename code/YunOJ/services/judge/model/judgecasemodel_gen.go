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
	judgeCaseFieldNames          = builder.RawFieldNames(&JudgeCase{})
	judgeCaseRows                = strings.Join(judgeCaseFieldNames, ",")
	judgeCaseRowsExpectAutoSet   = strings.Join(stringx.Remove(judgeCaseFieldNames, "`case_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	judgeCaseRowsWithPlaceHolder = strings.Join(stringx.Remove(judgeCaseFieldNames, "`case_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheJudgeCaseCaseIdPrefix = "cache:judgeCase:caseId:"
)

type (
	judgeCaseModel interface {
		Insert(ctx context.Context, data *JudgeCase) (sql.Result, error)
		FindOne(ctx context.Context, caseId int64) (*JudgeCase, error)
		Update(ctx context.Context, data *JudgeCase) error
		Delete(ctx context.Context, caseId int64) error
	}

	defaultJudgeCaseModel struct {
		sqlc.CachedConn
		table string
	}

	JudgeCase struct {
		CaseId     int64     `db:"case_id"`     // 测试用例id
		ProblemId  int64     `db:"problem_id"`  // 题目id
		Input      string    `db:"input"`       // 输入文件路径
		Output     string    `db:"output"`      // 输出文件路径
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newJudgeCaseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultJudgeCaseModel {
	return &defaultJudgeCaseModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`judge_case`",
	}
}

func (m *defaultJudgeCaseModel) Delete(ctx context.Context, caseId int64) error {
	judgeCaseCaseIdKey := fmt.Sprintf("%s%v", cacheJudgeCaseCaseIdPrefix, caseId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `case_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, caseId)
	}, judgeCaseCaseIdKey)
	return err
}

func (m *defaultJudgeCaseModel) FindOne(ctx context.Context, caseId int64) (*JudgeCase, error) {
	judgeCaseCaseIdKey := fmt.Sprintf("%s%v", cacheJudgeCaseCaseIdPrefix, caseId)
	var resp JudgeCase
	err := m.QueryRowCtx(ctx, &resp, judgeCaseCaseIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `case_id` = ? limit 1", judgeCaseRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, caseId)
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

func (m *defaultJudgeCaseModel) Insert(ctx context.Context, data *JudgeCase) (sql.Result, error) {
	judgeCaseCaseIdKey := fmt.Sprintf("%s%v", cacheJudgeCaseCaseIdPrefix, data.CaseId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, judgeCaseRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.ProblemId, data.Input, data.Output)
	}, judgeCaseCaseIdKey)
	return ret, err
}

func (m *defaultJudgeCaseModel) Update(ctx context.Context, data *JudgeCase) error {
	judgeCaseCaseIdKey := fmt.Sprintf("%s%v", cacheJudgeCaseCaseIdPrefix, data.CaseId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `case_id` = ?", m.table, judgeCaseRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.ProblemId, data.Input, data.Output, data.CaseId)
	}, judgeCaseCaseIdKey)
	return err
}

func (m *defaultJudgeCaseModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheJudgeCaseCaseIdPrefix, primary)
}

func (m *defaultJudgeCaseModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `case_id` = ? limit 1", judgeCaseRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultJudgeCaseModel) tableName() string {
	return m.table
}