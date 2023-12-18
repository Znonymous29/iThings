package hooks

import (
	"context"
	"github.com/i-Things/things/shared/def"
	"github.com/i-Things/things/shared/errors"
	"github.com/i-Things/things/shared/utils"
	"github.com/i-Things/things/src/vidsvr/internal/repo/relationDB"
	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type OnServerKeepaliveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnServerKeepaliveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnServerKeepaliveLogic {
	return &OnServerKeepaliveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnServerKeepaliveLogic) OnServerKeepalive(req *types.HooksApiServerKeepaliveReq) (resp *types.HooksApiResp, err error) {
	// todo: add your logic here and delete this line
	//hookactive中是保持在线状态  需要更新对应的数据库
	infoRepo := relationDB.NewVidmgrInfoRepo(l.ctx)
	vidmgrInfo, err := infoRepo.FindOneByFilter(l.ctx, relationDB.VidmgrFilter{
		VidmgrIDs: []string{req.MediaServerId},
	})

	if vidmgrInfo != nil {
		//update info
		vidmgrInfo.VidmgrStatus = def.DeviceStatusOnline
		vidmgrInfo.LastLogin = time.Now()

		err := infoRepo.Update(l.ctx, vidmgrInfo)
		if err != nil {
			er := errors.Fmt(err)
			l.Errorf("%s.rpc.ManageVidmgr req=%v err=%v", utils.FuncName(), req, er)
		}
	}
	return &types.HooksApiResp{
		Code: 0,
		Msg:  "success",
	}, nil
}
