package gateway

import (
	"context"
	"gitee.com/i-Things/share/ctxs"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"gitee.com/i-Things/things/service/apisvr/internal/logic"
	"gitee.com/i-Things/things/service/apisvr/internal/logic/things"
	"gitee.com/i-Things/things/service/dmsvr/pb/dm"

	"gitee.com/i-Things/things/service/apisvr/internal/svc"
	"gitee.com/i-Things/things/service/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctxs.WithDefaultRoot(ctx),
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.DeviceGateWayIndexReq) (resp *types.DeviceGateWayIndexResp, err error) {
	dmReq := &dm.DeviceGatewayIndexReq{
		Gateway: &dm.DeviceCore{
			ProductID:  req.GateWayProductID,
			DeviceName: req.GateWayDeviceName,
		},
		Page: logic.ToDmPageRpc(req.Page),
	}
	dmResp, err := l.svcCtx.DeviceM.DeviceGatewayIndex(l.ctx, dmReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.GetDeviceInfo req=%v err=%+v", utils.FuncName(), req, er)
		return nil, er
	}
	if dmResp.Total == 0 {
		return &types.DeviceGateWayIndexResp{
			List:  nil,
			Total: 0,
		}, nil
	}
	pis := make([]*types.DeviceInfo, 0, len(dmResp.List))
	ret, err := l.svcCtx.DeviceM.DeviceInfoIndex(l.ctx, &dm.DeviceInfoIndexReq{
		Devices: dmResp.List,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range ret.List {
		pi := things.InfoToApi(l.ctx, l.svcCtx, v, things.DeviceInfoWith{})
		pis = append(pis, pi)
	}
	return &types.DeviceGateWayIndexResp{
		Total: dmResp.Total,
		List:  pis,
	}, nil
}
