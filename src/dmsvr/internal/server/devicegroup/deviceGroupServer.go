// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/devicegroup"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type DeviceGroupServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedDeviceGroupServer
}

func NewDeviceGroupServer(svcCtx *svc.ServiceContext) *DeviceGroupServer {
	return &DeviceGroupServer{
		svcCtx: svcCtx,
	}
}

// 创建分组
func (s *DeviceGroupServer) GroupInfoCreate(ctx context.Context, in *dm.GroupInfoCreateReq) (*dm.Response, error) {
	l := devicegrouplogic.NewGroupInfoCreateLogic(ctx, s.svcCtx)
	return l.GroupInfoCreate(in)
}

// 获取分组信息列表
func (s *DeviceGroupServer) GroupInfoIndex(ctx context.Context, in *dm.GroupInfoIndexReq) (*dm.GroupInfoIndexResp, error) {
	l := devicegrouplogic.NewGroupInfoIndexLogic(ctx, s.svcCtx)
	return l.GroupInfoIndex(in)
}

// 获取分组信息详情
func (s *DeviceGroupServer) GroupInfoRead(ctx context.Context, in *dm.GroupInfoReadReq) (*dm.GroupInfo, error) {
	l := devicegrouplogic.NewGroupInfoReadLogic(ctx, s.svcCtx)
	return l.GroupInfoRead(in)
}

// 更新分组
func (s *DeviceGroupServer) GroupInfoUpdate(ctx context.Context, in *dm.GroupInfoUpdateReq) (*dm.Response, error) {
	l := devicegrouplogic.NewGroupInfoUpdateLogic(ctx, s.svcCtx)
	return l.GroupInfoUpdate(in)
}

// 删除分组
func (s *DeviceGroupServer) GroupInfoDelete(ctx context.Context, in *dm.GroupInfoDeleteReq) (*dm.Response, error) {
	l := devicegrouplogic.NewGroupInfoDeleteLogic(ctx, s.svcCtx)
	return l.GroupInfoDelete(in)
}

// 创建分组设备
func (s *DeviceGroupServer) GroupDeviceMultiCreate(ctx context.Context, in *dm.GroupDeviceMultiCreateReq) (*dm.Response, error) {
	l := devicegrouplogic.NewGroupDeviceMultiCreateLogic(ctx, s.svcCtx)
	return l.GroupDeviceMultiCreate(in)
}

// 获取分组设备信息列表
func (s *DeviceGroupServer) GroupDeviceIndex(ctx context.Context, in *dm.GroupDeviceIndexReq) (*dm.GroupDeviceIndexResp, error) {
	l := devicegrouplogic.NewGroupDeviceIndexLogic(ctx, s.svcCtx)
	return l.GroupDeviceIndex(in)
}

// 删除分组设备
func (s *DeviceGroupServer) GroupDeviceMultiDelete(ctx context.Context, in *dm.GroupDeviceMultiDeleteReq) (*dm.Response, error) {
	l := devicegrouplogic.NewGroupDeviceMultiDeleteLogic(ctx, s.svcCtx)
	return l.GroupDeviceMultiDelete(in)
}
