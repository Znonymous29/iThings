package deviceinteractlogic

import (
	"context"
	"gitee.com/i-Things/share/errors"
	"gitee.com/i-Things/share/utils"
	"github.com/i-Things/things/service/dmsvr/internal/repo/relationDB"
	"golang.org/x/sync/errgroup"
	"sync"

	"github.com/i-Things/things/service/dmsvr/internal/svc"
	"github.com/i-Things/things/service/dmsvr/pb/dm"

	"github.com/zeromicro/go-zero/core/logx"
)

type PropertyControlMultiSendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPropertyControlMultiSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PropertyControlMultiSendLogic {
	return &PropertyControlMultiSendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量调用设备属性
func (l *PropertyControlMultiSendLogic) PropertyControlMultiSend(in *dm.PropertyControlMultiSendReq) (*dm.PropertyControlMultiSendResp, error) {
	var list []*dm.PropertyControlSendMsg
	var err error
	if len(in.DeviceNames) != 0 {
		list, err = l.MultiSendOneProductProperty(in)
		if err != nil {
			return nil, err
		}
	} else {
		list, err = l.MultiSendMultiProductProperty(in)
		if err != nil {
			return nil, err
		}
	}
	return &dm.PropertyControlMultiSendResp{List: list}, nil
}

func (l *PropertyControlMultiSendLogic) MultiSendOneProductProperty(in *dm.PropertyControlMultiSendReq) ([]*dm.PropertyControlSendMsg, error) {
	list := make([]*dm.PropertyControlSendMsg, 0)
	sigSend := NewPropertyControlSendLogic(l.ctx, l.svcCtx)
	err := sigSend.initMsg(in.ProductID)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, v := range in.DeviceNames {
		wg.Add(1)
		go func(v string) {
			defer utils.Recover(l.ctx)
			defer wg.Done()
			ret, err := sigSend.PropertyControlSend(&dm.PropertyControlSendReq{
				ProductID:  in.ProductID,
				DeviceName: v,
				Data:       in.Data,
				IsAsync:    false,
			})

			if err != nil {
				myErr, _ := err.(*errors.CodeError)
				msg := &dm.PropertyControlSendMsg{
					DeviceName: v,
					SysMsg:     myErr.GetMsg(),
					SysCode:    myErr.Code,
				}
				mu.Lock()
				defer mu.Unlock()
				list = append(list, msg)
				return
			}

			msg := &dm.PropertyControlSendMsg{
				ProductID:  in.ProductID,
				DeviceName: v,
				SysCode:    errors.OK.Code,
				SysMsg:     errors.OK.GetMsg(),
				Code:       ret.Code,
				Msg:        ret.Msg,
				MsgToken:   ret.MsgToken,
			}
			mu.Lock()
			defer mu.Unlock()
			list = append(list, msg)
		}(v)
	}

	wg.Wait()
	return list, err
}

func (l *PropertyControlMultiSendLogic) MultiSendMultiProductProperty(in *dm.PropertyControlMultiSendReq) ([]*dm.PropertyControlSendMsg, error) {
	var productMap = map[string]map[string]struct{}{} //key是产品id,value是产品下的设备列表
	for _, v := range in.Devices {
		if productMap[v.ProductID] == nil {
			productMap[v.ProductID] = map[string]struct{}{v.DeviceName: {}}
		} else {
			productMap[v.ProductID][v.DeviceName] = struct{}{}
		}
	}
	if in.AreaID != 0 {
		dis, err := relationDB.NewDeviceInfoRepo(l.ctx).FindByFilter(l.ctx, relationDB.DeviceFilter{AreaIDs: []int64{in.AreaID}}, nil)
		if err != nil {
			return nil, err
		}
		for _, v := range dis {
			if productMap[v.ProductID] == nil {
				productMap[v.ProductID] = map[string]struct{}{v.DeviceName: {}}
			} else {
				productMap[v.ProductID][v.DeviceName] = struct{}{}
			}
		}
	}
	if in.GroupID != 0 {
		dis, err := relationDB.NewDeviceInfoRepo(l.ctx).FindByFilter(l.ctx, relationDB.DeviceFilter{GroupIDs: []int64{in.GroupID}}, nil)
		if err != nil {
			return nil, err
		}
		for _, v := range dis {
			if productMap[v.ProductID] == nil {
				productMap[v.ProductID] = map[string]struct{}{v.DeviceName: {}}
			} else {
				productMap[v.ProductID][v.DeviceName] = struct{}{}
			}
		}
	}
	var group errgroup.Group
	var newIn = dm.PropertyControlMultiSendReq{
		ShadowControl: in.ShadowControl,
		Data:          in.Data,
	}
	var mu sync.Mutex
	var list = []*dm.PropertyControlSendMsg{}
	for k, v := range productMap {
		in := newIn
		in.ProductID = k
		in.DeviceNames = utils.SetToSlice(v)
		group.Go(func() error {
			li, err := l.MultiSendOneProductProperty(&in)
			if err != nil {
				return err
			}
			mu.Lock()
			defer mu.Unlock()
			list = append(list, li...)
			return nil
		})
	}
	err := group.Wait()
	if err != nil {
		return nil, err
	}
	return list, err
}
