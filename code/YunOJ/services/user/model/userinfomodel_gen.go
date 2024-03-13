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
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`userid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`userid`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserInfoUseridPrefix   = "cache:userInfo:userid:"
	cacheUserInfoPhonePrefix    = "cache:userInfo:phone:"
	cacheUserInfoUsernamePrefix = "cache:userInfo:username:"
)

type (
	userInfoModel interface {
		Insert(ctx context.Context, data *UserInfo) (sql.Result, error)
		FindOne(ctx context.Context, userid int64) (*UserInfo, error)
		FindOneByPhone(ctx context.Context, phone string) (*UserInfo, error)
		FindOneByUsername(ctx context.Context, username string) (*UserInfo, error)
		FindOneByEmail(ctx context.Context, email string) (*UserInfo, error)
		Update(ctx context.Context, data *UserInfo) error
		Delete(ctx context.Context, userid int64) error
	}

	defaultUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	UserInfo struct {
		Userid    int64     `db:"userid"`     // 用户ID
		Username  string    `db:"username"`   // 用户名称
		Email     string    `db:"email"`      // 电子邮件地址
		Phone     string    `db:"phone"`      // 手机号码
		Gender    int64     `db:"gender"`     // 性别：1-男，2-女
		UserType  int64     `db:"user_type"`  // 用户类型：1-教练，2-队员
		Password  string    `db:"password"`   // 密码
		Avatar    string    `db:"avatar"`     // 头像
		Status    int64     `db:"status"`     // 用户状态：1-激活，2-禁用，3-待审核
		CreatedAt time.Time `db:"created_at"` // 记录创建时间
		UpdatedAt time.Time `db:"updated_at"` // 记录更新时间
	}
)

func newUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserInfoModel {
	return &defaultUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`userInfo`",
	}
}

func (m *defaultUserInfoModel) Delete(ctx context.Context, userid int64) error {
	data, err := m.FindOne(ctx, userid)
	if err != nil {
		return err
	}

	userInfoPhoneKey := fmt.Sprintf("%s%v", cacheUserInfoPhonePrefix, data.Phone)
	userInfoUseridKey := fmt.Sprintf("%s%v", cacheUserInfoUseridPrefix, userid)
	userInfoUsernameKey := fmt.Sprintf("%s%v", cacheUserInfoUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `userid` = ?", m.table)
		return conn.ExecCtx(ctx, query, userid)
	}, userInfoPhoneKey, userInfoUseridKey, userInfoUsernameKey)
	return err
}

func (m *defaultUserInfoModel) FindOne(ctx context.Context, userid int64) (*UserInfo, error) {
	userInfoUseridKey := fmt.Sprintf("%s%v", cacheUserInfoUseridPrefix, userid)
	var resp UserInfo
	err := m.QueryRowCtx(ctx, &resp, userInfoUseridKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", userInfoRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, userid)
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

func (m *defaultUserInfoModel) FindOneByPhone(ctx context.Context, phone string) (*UserInfo, error) {
	userInfoPhoneKey := fmt.Sprintf("%s%v", cacheUserInfoPhonePrefix, phone)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, userInfoPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `phone` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Userid, nil
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

func (m *defaultUserInfoModel) FindOneByUsername(ctx context.Context, username string) (*UserInfo, error) {
	userInfoUsernameKey := fmt.Sprintf("%s%v", cacheUserInfoUsernamePrefix, username)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, userInfoUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `username` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
			return nil, err
		}
		return resp.Userid, nil
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

func (m *defaultUserInfoModel) FindOneByEmail(ctx context.Context, email string) (*UserInfo, error) {
	userInfoUsernameKey := fmt.Sprintf("%s%v", cacheUserInfoUsernamePrefix, email)
	var resp UserInfo
	err := m.QueryRowIndexCtx(ctx, &resp, userInfoUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `email` = ? limit 1", userInfoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Userid, nil
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

func (m *defaultUserInfoModel) Insert(ctx context.Context, data *UserInfo) (sql.Result, error) {
	userInfoPhoneKey := fmt.Sprintf("%s%v", cacheUserInfoPhonePrefix, data.Phone)
	userInfoUseridKey := fmt.Sprintf("%s%v", cacheUserInfoUseridPrefix, data.Userid)
	userInfoUsernameKey := fmt.Sprintf("%s%v", cacheUserInfoUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Email, data.Phone, data.Gender, data.UserType, data.Password, data.Avatar, data.Status)
	}, userInfoPhoneKey, userInfoUseridKey, userInfoUsernameKey)
	return ret, err
}

func (m *defaultUserInfoModel) Update(ctx context.Context, newData *UserInfo) error {
	data, err := m.FindOne(ctx, newData.Userid)
	if err != nil {
		return err
	}

	userInfoPhoneKey := fmt.Sprintf("%s%v", cacheUserInfoPhonePrefix, data.Phone)
	userInfoUseridKey := fmt.Sprintf("%s%v", cacheUserInfoUseridPrefix, data.Userid)
	userInfoUsernameKey := fmt.Sprintf("%s%v", cacheUserInfoUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `userid` = ?", m.table, userInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Username, newData.Email, newData.Phone, newData.Gender, newData.UserType, newData.Password, newData.Avatar, newData.Status, newData.Userid)
	}, userInfoPhoneKey, userInfoUseridKey, userInfoUsernameKey)
	return err
}

func (m *defaultUserInfoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserInfoUseridPrefix, primary)
}

func (m *defaultUserInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `userid` = ? limit 1", userInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserInfoModel) tableName() string {
	return m.table
}
