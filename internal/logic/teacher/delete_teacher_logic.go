package teacher

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/client/school"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewDeleteTeacherLogic(r *http.Request, svcCtx *svc.ServiceContext) *DeleteTeacherLogic {
	return &DeleteTeacherLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *DeleteTeacherLogic) DeleteTeacher(req *types.UUIDReq) (resp *types.BaseMsgResp, err error) {
	result, err := l.svcCtx.SchoolRpc.DeleteTeacher(l.ctx, &example.UUIDReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, result.Msg)}, nil
}
