// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysApiAuthFieldNames          = builder.RawFieldNames(&SysApiAuth{})
	sysApiAuthRows                = strings.Join(sysApiAuthFieldNames, ",")
	sysApiAuthRowsExpectAutoSet   = strings.Join(stringx.Remove(sysApiAuthFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), ",")
	sysApiAuthRowsWithPlaceHolder = strings.Join(stringx.Remove(sysApiAuthFieldNames, "`id`", "`createdTime`", "`deletedTime`", "`updatedTime`"), "=?,") + "=?"
)

type (
	sysApiAuthModel interface {
		Insert(ctx context.Context, data *SysApiAuth) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysApiAuth, error)
		FindOneByV0V1(ctx context.Context, v0 string, v1 string) (*SysApiAuth, error)
		Update(ctx context.Context, data *SysApiAuth) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysApiAuthModel struct {
		conn  sqlx.SqlConn
		table string
	}

	SysApiAuth struct {
		Id    int64  `db:"id"`     // 编号
		PType string `db:"p_type"` // 策略类型，即策略的分类，例如"p"表示主体（provider）访问资源（resource）的许可权，"g"表示主体（provider）之间的关系访问控制
		V0    string `db:"v0"`     // 策略中的第一个参数，通常用于表示资源的归属范围（即限制访问的对象），例如资源所属的机构、部门、业务线、地域等
		V1    string `db:"v1"`     // 策略中的第二个参数，通常用于表示主体（provider），即需要访问资源的用户或者服务
		V2    string `db:"v2"`     // 策略中的第三个参数，通常用于表示资源（resource），即需要进行访问的对象
		V3    string `db:"v3"`     // 策略中的第四个参数，通常用于表示访问操作（permission），例如 “read”, “write”, “execute” 等
		V4    string `db:"v4"`     // 策略中的第五个参数，通常用于表示资源的类型（object type），例如表示是文件或者数据库表等
		V5    string `db:"v5"`     // 策略中的第六个参数，通常用于表示扩展信息，例如 IP 地址、端口号等
	}
)

func newSysApiAuthModel(conn sqlx.SqlConn) *defaultSysApiAuthModel {
	return &defaultSysApiAuthModel{
		conn:  conn,
		table: "`sys_api_auth`",
	}
}

func (m *defaultSysApiAuthModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultSysApiAuthModel) FindOne(ctx context.Context, id int64) (*SysApiAuth, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysApiAuthRows, m.table)
	var resp SysApiAuth
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiAuthModel) FindOneByV0V1(ctx context.Context, v0 string, v1 string) (*SysApiAuth, error) {
	var resp SysApiAuth
	query := fmt.Sprintf("select %s from %s where `v0` = ? and `v1` = ? limit 1", sysApiAuthRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, v0, v1)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysApiAuthModel) Insert(ctx context.Context, data *SysApiAuth) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysApiAuthRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.PType, data.V0, data.V1, data.V2, data.V3, data.V4, data.V5)
	return ret, err
}

func (m *defaultSysApiAuthModel) Update(ctx context.Context, newData *SysApiAuth) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysApiAuthRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.PType, newData.V0, newData.V1, newData.V2, newData.V3, newData.V4, newData.V5, newData.Id)
	return err
}

func (m *defaultSysApiAuthModel) tableName() string {
	return m.table
}
