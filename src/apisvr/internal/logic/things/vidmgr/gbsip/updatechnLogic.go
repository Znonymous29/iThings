package gbsip

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/i-Things/things/src/apisvr/internal/svc"
	"github.com/i-Things/things/src/apisvr/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatechnLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatechnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatechnLogic {
	return &UpdatechnLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdatechnLogic) Updatechn(req *types.VidmgrSipUpdateChnReq) error {
	// todo: add your logic here and delete this line
	if req.ChannelID == "" {
		return errors.MediaGbsipUpdateError
	}

	vidReq := &vid.VidmgrGbsipChannelUpdate{
		ChannelID:  req.ChannelID,
		Memo:       req.Memo,
		StreamType: req.StreamType,
		Url:        req.URL,
	}
	jsonStr, _ := json.Marshal(req)
	fmt.Println("airgens Updatedev:", string(jsonStr))
	_, err := l.svcCtx.VidmgrG.VidmgrGbsipChannelUpdate(l.ctx, vidReq)
	if err != nil {
		er := errors.Fmt(err)
		l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
		return er
	}
	return nil
}
