package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudentLogic) UpdateStudent(req *types.StudentInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.UpdateStudent(l.ctx,
		&example.StudentInfo{
			Id:            req.Id,
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
