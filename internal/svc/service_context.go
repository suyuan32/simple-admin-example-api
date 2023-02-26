package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-example-rpc/exampleclient"

	"github.com/suyuan32/simple-admin-example-api/internal/config"
	"github.com/suyuan32/simple-admin-example-api/internal/middleware"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	Casbin     *casbin.Enforcer
	Authority  rest.Middleware
	Trans      *i18n.Translator
	ExampleRpc exampleclient.Example
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbin(c.DatabaseConf.Type, c.DatabaseConf.GetDSN())

	trans := &i18n.Translator{}
	trans.NewBundle(i18n.LocaleFS)
	trans.NewTranslator()

	return &ServiceContext{
		Config:     c,
		Authority:  middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:      trans,
		ExampleRpc: exampleclient.NewExample(zrpc.MustNewClient(c.ExampleRpc)),
	}
}
