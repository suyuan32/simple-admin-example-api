package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTeacherLogic {
	return &CreateTeacherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTeacherLogic) CreateTeacher(req *types.TeacherInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.CreateTeacher(l.ctx,
		&example.TeacherInfo{
			Name:          req.Name,
			Age:           req.Age,
			AgeInt32:      req.AgeInt32,
			AgeInt64:      req.AgeInt64,
			AgeUint:       req.AgeUint,
			AgeUint32:     req.AgeUint32,
			AgeUint64:     req.AgeUint64,
			WeightFloat:   req.WeightFloat,
			WeightFloat32: req.WeightFloat32,
			ClassId:       req.ClassId,
			EnrollAt:      req.EnrollAt,
			StatusBool:    req.StatusBool,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
