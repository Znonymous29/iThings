package deviceinteractlogic

import (
	"gitee.com/i-Things/things/service/dmsvr/internal/domain/serverDo"
	"gitee.com/i-Things/things/service/dmsvr/pb/dm"
)

func ToSendOptionDo(in *dm.SendOption) *serverDo.SendOption {
	if in == nil {
		return nil
	}
	return &serverDo.SendOption{
		TimeoutToFail:  in.TimeoutToFail,
		RequestTimeout: in.RequestTimeout,
		RetryInterval:  in.RetryInterval,
	}
}
