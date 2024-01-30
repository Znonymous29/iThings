// Code generated by goctl. DO NOT EDIT.
// Source: ud.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/udsvr/internal/logic/userdevice"
	"github.com/i-Things/things/src/udsvr/internal/svc"
	"github.com/i-Things/things/src/udsvr/pb/ud"
)

type UserDeviceServer struct {
	svcCtx *svc.ServiceContext
	ud.UnimplementedUserDeviceServer
}

func NewUserDeviceServer(svcCtx *svc.ServiceContext) *UserDeviceServer {
	return &UserDeviceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserDeviceServer) UserCollectDeviceMultiCreate(ctx context.Context, in *ud.UserCollectDeviceSave) (*ud.Empty, error) {
	l := userdevicelogic.NewUserCollectDeviceMultiCreateLogic(ctx, s.svcCtx)
	return l.UserCollectDeviceMultiCreate(in)
}

func (s *UserDeviceServer) UserCollectDeviceMultiDelete(ctx context.Context, in *ud.UserCollectDeviceSave) (*ud.Empty, error) {
	l := userdevicelogic.NewUserCollectDeviceMultiDeleteLogic(ctx, s.svcCtx)
	return l.UserCollectDeviceMultiDelete(in)
}

func (s *UserDeviceServer) UserCollectDeviceIndex(ctx context.Context, in *ud.Empty) (*ud.UserCollectDeviceSave, error) {
	l := userdevicelogic.NewUserCollectDeviceIndexLogic(ctx, s.svcCtx)
	return l.UserCollectDeviceIndex(in)
}
