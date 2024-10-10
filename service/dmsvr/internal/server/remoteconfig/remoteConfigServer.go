// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic/remoteconfig"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
)

type RemoteConfigServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedRemoteConfigServer
}

func NewRemoteConfigServer(svcCtx *svc.ServiceContext) *RemoteConfigServer {
	return &RemoteConfigServer{
		svcCtx: svcCtx,
	}
}

func (s *RemoteConfigServer) RemoteConfigCreate(ctx context.Context, in *dm.RemoteConfigCreateReq) (*dm.Empty, error) {
	l := remoteconfiglogic.NewRemoteConfigCreateLogic(ctx, s.svcCtx)
	return l.RemoteConfigCreate(in)
}

func (s *RemoteConfigServer) RemoteConfigIndex(ctx context.Context, in *dm.RemoteConfigIndexReq) (*dm.RemoteConfigIndexResp, error) {
	l := remoteconfiglogic.NewRemoteConfigIndexLogic(ctx, s.svcCtx)
	return l.RemoteConfigIndex(in)
}

func (s *RemoteConfigServer) RemoteConfigPushAll(ctx context.Context, in *dm.RemoteConfigPushAllReq) (*dm.Empty, error) {
	l := remoteconfiglogic.NewRemoteConfigPushAllLogic(ctx, s.svcCtx)
	return l.RemoteConfigPushAll(in)
}

func (s *RemoteConfigServer) RemoteConfigLastRead(ctx context.Context, in *dm.RemoteConfigLastReadReq) (*dm.RemoteConfigLastReadResp, error) {
	l := remoteconfiglogic.NewRemoteConfigLastReadLogic(ctx, s.svcCtx)
	return l.RemoteConfigLastRead(in)
}
