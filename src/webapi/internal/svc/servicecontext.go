package svc

import (
	"gitee.com/godLei6/things/shared/verify"
	"gitee.com/godLei6/things/src/usersvr/model"
	"gitee.com/godLei6/things/src/usersvr/userclient"
	"gitee.com/godLei6/things/src/webapi/internal/config"
	"gitee.com/godLei6/things/src/webapi/internal/middleware"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/rest"
	"github.com/tal-tech/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config        config.Config
	CheckToken    rest.Middleware
	Record        rest.Middleware
	UserInfoModel model.UserInfoModel
	UserCoreModel model.UserCoreModel
	UserRpc       userclient.User
	Captcha       *verify.Captcha
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	ui := model.NewUserInfoModel(conn, c.CacheRedis)
	uc := model.NewUserCoreModel(conn, c.CacheRedis)
	ur := userclient.NewUser(zrpc.MustNewClient(c.UserRpc))
	captcha := verify.NewCaptcha(c.ImgHeight, c.ImgWidth, c.KeyLong, c.CacheRedis, time.Duration(c.KeepTime)*time.Second)
	return &ServiceContext{
		Config:        c,
		CheckToken:    middleware.NewCheckTokenMiddleware(ur).Handle,
		Record:        middleware.NewRecordMiddleware().Handle,
		UserInfoModel: ui,
		UserCoreModel: uc,
		UserRpc:       ur,
		Captcha:       captcha,
	}
}
