// Code generated by goctl. DO NOT EDIT.
// Source: sys.proto

package common

import (
	"context"

	"github.com/i-Things/things/src/syssvr/internal/svc"
	"github.com/i-Things/things/src/syssvr/pb/sys"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ApiCreateReq          = sys.ApiCreateReq
	ApiData               = sys.ApiData
	ApiDeleteReq          = sys.ApiDeleteReq
	ApiIndexReq           = sys.ApiIndexReq
	ApiIndexResp          = sys.ApiIndexResp
	ApiUpdateReq          = sys.ApiUpdateReq
	AuthApiIndexReq       = sys.AuthApiIndexReq
	AuthApiIndexResp      = sys.AuthApiIndexResp
	AuthApiInfo           = sys.AuthApiInfo
	AuthApiMultiUpdateReq = sys.AuthApiMultiUpdateReq
	CheckAuthReq          = sys.CheckAuthReq
	ConfigResp            = sys.ConfigResp
	DateRange             = sys.DateRange
	JwtToken              = sys.JwtToken
	LoginLogCreateReq     = sys.LoginLogCreateReq
	LoginLogIndexData     = sys.LoginLogIndexData
	LoginLogIndexReq      = sys.LoginLogIndexReq
	LoginLogIndexResp     = sys.LoginLogIndexResp
	Map                   = sys.Map
	MenuCreateReq         = sys.MenuCreateReq
	MenuData              = sys.MenuData
	MenuDeleteReq         = sys.MenuDeleteReq
	MenuIndexReq          = sys.MenuIndexReq
	MenuIndexResp         = sys.MenuIndexResp
	MenuUpdateReq         = sys.MenuUpdateReq
	OperLogCreateReq      = sys.OperLogCreateReq
	OperLogIndexData      = sys.OperLogIndexData
	OperLogIndexReq       = sys.OperLogIndexReq
	OperLogIndexResp      = sys.OperLogIndexResp
	PageInfo              = sys.PageInfo
	PageInfo_OrderBy      = sys.PageInfo_OrderBy
	Response              = sys.Response
	RoleCreateReq         = sys.RoleCreateReq
	RoleDeleteReq         = sys.RoleDeleteReq
	RoleIndexData         = sys.RoleIndexData
	RoleIndexReq          = sys.RoleIndexReq
	RoleIndexResp         = sys.RoleIndexResp
	RoleMenuUpdateReq     = sys.RoleMenuUpdateReq
	RoleUpdateReq         = sys.RoleUpdateReq
	UserCheckTokenReq     = sys.UserCheckTokenReq
	UserCheckTokenResp    = sys.UserCheckTokenResp
	UserCreateResp        = sys.UserCreateResp
	UserDeleteReq         = sys.UserDeleteReq
	UserIndexReq          = sys.UserIndexReq
	UserIndexResp         = sys.UserIndexResp
	UserInfo              = sys.UserInfo
	UserLoginReq          = sys.UserLoginReq
	UserLoginResp         = sys.UserLoginResp
	UserReadReq           = sys.UserReadReq

	Common interface {
		Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error)
	}

	defaultCommon struct {
		cli zrpc.Client
	}

	directCommon struct {
		svcCtx *svc.ServiceContext
		svr    sys.CommonServer
	}
)

func NewCommon(cli zrpc.Client) Common {
	return &defaultCommon{
		cli: cli,
	}
}

func NewDirectCommon(svcCtx *svc.ServiceContext, svr sys.CommonServer) Common {
	return &directCommon{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

func (m *defaultCommon) Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error) {
	client := sys.NewCommonClient(m.cli.Conn())
	return client.Config(ctx, in, opts...)
}

func (d *directCommon) Config(ctx context.Context, in *Response, opts ...grpc.CallOption) (*ConfigResp, error) {
	return d.svr.Config(ctx, in)
}
