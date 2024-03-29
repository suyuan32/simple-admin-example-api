package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherLogic {
	return &UpdateTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTeacherLogic) UpdateTeacher(req *types.TeacherInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.UpdateTeacher(l.ctx,
		&example.TeacherInfo{
			Id:   req.Id,
			Name: req.Name,
			Age:  req.Age,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
