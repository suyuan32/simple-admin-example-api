package svc

import (
	exampleClient "github.com/suyuan32/simple-admin-example-rpc/client/example"
	schoolClient "github.com/suyuan32/simple-admin-example-rpc/client/school"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/suyuan32/simple-admin-example-api/internal/config"
	"github.com/suyuan32/simple-admin-example-api/internal/middleware"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	ExampleRpc exampleClient.Example
	SchoolRpc  schoolClient.School
	Casbin     *casbin.Enforcer
	Authority  rest.Middleware
	Trans      *i18n.Translator
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := redis.MustNewRedis(c.RedisConf)

	cbn := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n.LocaleFS)

	return &ServiceContext{
		Config:     c,
		Authority:  middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:      trans,
		ExampleRpc: exampleClient.NewExample(zrpc.MustNewClient(c.ExampleRpc)),
		SchoolRpc:  schoolClient.NewSchool(zrpc.MustNewClient(c.SchoolRpc)),
	}
}
