package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentLogic {
	return &CreateStudentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStudentLogic) CreateStudent(req *types.StudentInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.CreateStudent(l.ctx,
		&example.StudentInfo{
			Id:            req.Id,
			Name:          req.Name,
			Age:           req.Age,
			AgeInt8:       req.AgeInt8,
			AgeUint8:      req.AgeUint8,
			AgeInt16:      req.AgeInt16,
			AgeUint16:     req.AgeUint16,
			AgeInt32:      req.AgeInt32,
			AgeUint32:     req.AgeUint32,
			AgeInt64:      req.AgeInt64,
			AgeUint64:     req.AgeUint64,
			AgeInt:        req.AgeInt,
			AgeUint:       req.AgeUint,
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
