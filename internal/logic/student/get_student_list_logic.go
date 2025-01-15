package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.StudentListReq) (resp *types.StudentListResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetStudentList(l.ctx,
		&example.StudentListReq{
			Page:          req.Page,
			PageSize:      req.PageSize,
			Status:        req.Status,
			Name:          req.Name,
			Age:           req.Age,
			Address:       req.Address,
			Score:         req.Score,
			Weight:        req.Weight,
			Healthy:       req.Healthy,
			Code:          req.Code,
			IdentifyId:    req.IdentifyId,
			Height:        req.Height,
			ExpiredAt:     req.ExpiredAt,
			StudentNumber: req.StudentNumber,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.StudentListResp{}
	resp.Msg = "successful"
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.StudentInfo{
				BaseUUIDInfo: types.BaseUUIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Status:        v.Status,
				Name:          v.Name,
				Age:           v.Age,
				Address:       v.Address,
				Score:         v.Score,
				Weight:        v.Weight,
				Healthy:       v.Healthy,
				Code:          v.Code,
				IdentifyId:    v.IdentifyId,
				Height:        v.Height,
				ExpiredAt:     v.ExpiredAt,
				StudentNumber: v.StudentNumber,
			})
	}
	return resp, nil
}
