package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentByIdLogic) GetStudentById(req *types.IDReq) (resp *types.StudentInfoResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetStudentById(l.ctx, &example.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.StudentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.StudentInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name:          data.Name,
			Age:           data.Age,
			AgeInt8:       data.AgeInt8,
			AgeUint8:      data.AgeUint8,
			AgeInt16:      data.AgeInt16,
			AgeUint16:     data.AgeUint16,
			AgeInt32:      data.AgeInt32,
			AgeUint32:     data.AgeUint32,
			AgeInt64:      data.AgeInt64,
			AgeUint64:     data.AgeUint64,
			AgeInt:        data.AgeInt,
			AgeUint:       data.AgeUint,
			WeightFloat:   data.WeightFloat,
			WeightFloat32: data.WeightFloat32,
			ClassId:       data.ClassId,
			EnrollAt:      data.EnrollAt,
			StatusBool:    data.StatusBool,
		},
	}, nil
}
