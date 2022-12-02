package student

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewBatchDeleteStudentLogic(r *http.Request, svcCtx *svc.ServiceContext) *BatchDeleteStudentLogic {
	return &BatchDeleteStudentLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *BatchDeleteStudentLogic) BatchDeleteStudent(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.ExampleRpc.BatchDeleteStudent(l.ctx, &example.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
