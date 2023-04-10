package innerSub

import (
	"context"
	"fmt"
	"github.com/i-Things/things/shared/devices"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/ddsvr/internal/domain/script"
	"github.com/i-Things/things/src/ddsvr/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type InnerSubServer struct {
	svcCtx *svc.ServiceContext
	logx.Logger
	ctx context.Context
}

func NewInnerSubServer(svcCtx *svc.ServiceContext, ctx context.Context) *InnerSubServer {
	return &InnerSubServer{
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
	}
}

func (s *InnerSubServer) PublishToDev(info *devices.InnerPublish) error {
	var finalPayload = info.Payload
	topic := fmt.Sprintf("%s/down/%s/%s/%s", "$"+info.Handle, info.Type, info.ProductID, info.DeviceName)

	f, err := s.svcCtx.Script.GetProtoFunc(s.ctx, info.ProductID, script.ConvertTypeProtoToRow, info.Handle, info.Type)
	if err != nil {
		s.Errorf("%s.GetProtoFunc info:%#v err:%v", utils.FuncName(), info, err)
		return err
	}
	if f != nil { //如果写了自定义函数
		finalPayload, err = f(info.Payload)
		if err != nil {
			s.Errorf("%s.Transform info:%#v err:%v", utils.FuncName(), info, err)
			return err
		}
		s.Infof("%s.transform success before:%#v after:%#v", utils.FuncName(), info.Payload, finalPayload)
		topic = fmt.Sprintf("%s/down/%s/%s/%s/%s",
			"$"+info.Handle, info.Type, script.CustomType, info.ProductID, info.DeviceName)
	}

	return s.svcCtx.PubDev.Publish(s.ctx, topic, finalPayload)
}
