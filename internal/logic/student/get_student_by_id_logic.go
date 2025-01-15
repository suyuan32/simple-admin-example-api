package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

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

func (l *GetStudentByIdLogic) GetStudentById(req *types.UUIDReq) (resp *types.StudentInfoResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetStudentById(l.ctx, &example.UUIDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.StudentInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  "successful",
		},
		Data: types.StudentInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Status:        data.Status,
			Name:          data.Name,
			Age:           data.Age,
			Address:       data.Address,
			Score:         data.Score,
			Weight:        data.Weight,
			Healthy:       data.Healthy,
			Code:          data.Code,
			IdentifyId:    data.IdentifyId,
			Height:        data.Height,
			ExpiredAt:     data.ExpiredAt,
			StudentNumber: data.StudentNumber,
		},
	}, nil
}
