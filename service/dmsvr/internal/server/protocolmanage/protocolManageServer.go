// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.1
// Source: dm.proto

package server

import (
	"context"

	"gitee.com/unitedrhino/things/service/dmsvr/internal/logic/protocolmanage"
	"gitee.com/unitedrhino/things/service/dmsvr/internal/svc"
	"gitee.com/unitedrhino/things/service/dmsvr/pb/dm"
)

type ProtocolManageServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedProtocolManageServer
}

func NewProtocolManageServer(svcCtx *svc.ServiceContext) *ProtocolManageServer {
	return &ProtocolManageServer{
		svcCtx: svcCtx,
	}
}

// 协议列表
func (s *ProtocolManageServer) ProtocolInfoIndex(ctx context.Context, in *dm.ProtocolInfoIndexReq) (*dm.ProtocolInfoIndexResp, error) {
	l := protocolmanagelogic.NewProtocolInfoIndexLogic(ctx, s.svcCtx)
	return l.ProtocolInfoIndex(in)
}

// 协议详情
func (s *ProtocolManageServer) ProtocolInfoRead(ctx context.Context, in *dm.WithIDCode) (*dm.ProtocolInfo, error) {
	l := protocolmanagelogic.NewProtocolInfoReadLogic(ctx, s.svcCtx)
	return l.ProtocolInfoRead(in)
}

// 协议创建
func (s *ProtocolManageServer) ProtocolInfoCreate(ctx context.Context, in *dm.ProtocolInfo) (*dm.WithID, error) {
	l := protocolmanagelogic.NewProtocolInfoCreateLogic(ctx, s.svcCtx)
	return l.ProtocolInfoCreate(in)
}

// 协议更新
func (s *ProtocolManageServer) ProtocolInfoUpdate(ctx context.Context, in *dm.ProtocolInfo) (*dm.Empty, error) {
	l := protocolmanagelogic.NewProtocolInfoUpdateLogic(ctx, s.svcCtx)
	return l.ProtocolInfoUpdate(in)
}

// 协议删除
func (s *ProtocolManageServer) ProtocolInfoDelete(ctx context.Context, in *dm.WithID) (*dm.Empty, error) {
	l := protocolmanagelogic.NewProtocolInfoDeleteLogic(ctx, s.svcCtx)
	return l.ProtocolInfoDelete(in)
}
