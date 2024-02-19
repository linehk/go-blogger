// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	localeFieldNames          = builder.RawFieldNames(&Locale{}, true)
	localeRows                = strings.Join(localeFieldNames, ",")
	localeRowsExpectAutoSet   = strings.Join(stringx.Remove(localeFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	localeRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(localeFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicLocaleIdPrefix          = "cache:public:locale:id:"
	cachePublicLocaleAppUserUuidPrefix = "cache:public:locale:appUserUuid:"
	cachePublicLocaleUuidPrefix        = "cache:public:locale:uuid:"
)

type (
	localeModel interface {
		Insert(ctx context.Context, data *Locale) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Locale, error)
		FindOneByAppUserUuid(ctx context.Context, appUserUuid string) (*Locale, error)
		FindOneByUuid(ctx context.Context, uuid string) (*Locale, error)
		Update(ctx context.Context, data *Locale) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLocaleModel struct {
		sqlc.CachedConn
		table string
	}

	Locale struct {
		Id          int64          `db:"id"`
		Uuid        string         `db:"uuid"`
		AppUserUuid string         `db:"app_user_uuid"`
		Language    sql.NullString `db:"language"`
		Country     sql.NullString `db:"country"`
		Variant     sql.NullString `db:"variant"`
	}
)

func newLocaleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultLocaleModel {
	return &defaultLocaleModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."locale"`,
	}
}

func (m *defaultLocaleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicLocaleAppUserUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleAppUserUuidPrefix, data.AppUserUuid)
	publicLocaleIdKey := fmt.Sprintf("%s%v", cachePublicLocaleIdPrefix, id)
	publicLocaleUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicLocaleAppUserUuidKey, publicLocaleIdKey, publicLocaleUuidKey)
	return err
}

func (m *defaultLocaleModel) FindOne(ctx context.Context, id int64) (*Locale, error) {
	publicLocaleIdKey := fmt.Sprintf("%s%v", cachePublicLocaleIdPrefix, id)
	var resp Locale
	err := m.QueryRowCtx(ctx, &resp, publicLocaleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", localeRows, m.table)
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

func (m *defaultLocaleModel) FindOneByAppUserUuid(ctx context.Context, appUserUuid string) (*Locale, error) {
	publicLocaleAppUserUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleAppUserUuidPrefix, appUserUuid)
	var resp Locale
	err := m.QueryRowIndexCtx(ctx, &resp, publicLocaleAppUserUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where app_user_uuid = $1 limit 1", localeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, appUserUuid); err != nil {
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

func (m *defaultLocaleModel) FindOneByUuid(ctx context.Context, uuid string) (*Locale, error) {
	publicLocaleUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleUuidPrefix, uuid)
	var resp Locale
	err := m.QueryRowIndexCtx(ctx, &resp, publicLocaleUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where uuid = $1 limit 1", localeRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uuid); err != nil {
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

func (m *defaultLocaleModel) Insert(ctx context.Context, data *Locale) (sql.Result, error) {
	publicLocaleAppUserUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleAppUserUuidPrefix, data.AppUserUuid)
	publicLocaleIdKey := fmt.Sprintf("%s%v", cachePublicLocaleIdPrefix, data.Id)
	publicLocaleUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5)", m.table, localeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.AppUserUuid, data.Language, data.Country, data.Variant)
	}, publicLocaleAppUserUuidKey, publicLocaleIdKey, publicLocaleUuidKey)
	return ret, err
}

func (m *defaultLocaleModel) Update(ctx context.Context, newData *Locale) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicLocaleAppUserUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleAppUserUuidPrefix, data.AppUserUuid)
	publicLocaleIdKey := fmt.Sprintf("%s%v", cachePublicLocaleIdPrefix, data.Id)
	publicLocaleUuidKey := fmt.Sprintf("%s%v", cachePublicLocaleUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, localeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Uuid, newData.AppUserUuid, newData.Language, newData.Country, newData.Variant)
	}, publicLocaleAppUserUuidKey, publicLocaleIdKey, publicLocaleUuidKey)
	return err
}

func (m *defaultLocaleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicLocaleIdPrefix, primary)
}

func (m *defaultLocaleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", localeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLocaleModel) tableName() string {
	return m.table
}
