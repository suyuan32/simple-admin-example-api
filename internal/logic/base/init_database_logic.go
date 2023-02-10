package base

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewInitDatabaseLogic(r *http.Request, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line

	return
}
