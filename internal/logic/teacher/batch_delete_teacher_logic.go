package teacher

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewBatchDeleteTeacherLogic(r *http.Request, svcCtx *svc.ServiceContext) *BatchDeleteTeacherLogic {
	return &BatchDeleteTeacherLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *BatchDeleteTeacherLogic) BatchDeleteTeacher(req *types.UUIDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.SchoolRpc.BatchDeleteTeacher(l.ctx, &example.UUIDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
