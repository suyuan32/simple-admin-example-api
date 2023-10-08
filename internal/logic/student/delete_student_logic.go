package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentLogic {
	return &DeleteStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStudentLogic) DeleteStudent(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.DeleteStudent(l.ctx, &example.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
