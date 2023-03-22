package log

import (
	"context"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/apisvr/internal/logic"
	"github.com/i-Things/things/src/rulesvr/pb/rule"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

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
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index(req *types.AlarmLogIndexReq) (resp *types.AlarmLogIndexResp, err error) {
	ret, err := l.svcCtx.Alarm.AlarmLogIndex(l.ctx, &rule.AlarmLogIndexReq{
		Page:      logic.ToRulePageRpc(req.Page),
		TimeRange: logic.ToRuleTimeRangeRpc(req.TimeRange),
	})
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.AlarmLogIndex req=%v err=%v", utils.FuncName(), req, er)
		return nil, er
	}
	pis := make([]*types.AlarmLog, 0, len(ret.List))
	for _, v := range ret.List {
		pi := &types.AlarmLog{
			ID:          v.Id,
			Desc:        v.Desc,
			CreatedTime: v.CreatedTime,
			AlarmID:     v.AlarmID,
			Serial:      v.Serial,
			SceneName:   v.SceneName,
			SceneID:     v.SceneID,
		}
		pis = append(pis, pi)
	}
	return &types.AlarmLogIndexResp{
		Total: ret.Total,
		List:  pis,
		Num:   int64(len(pis)),
	}, nil
}
