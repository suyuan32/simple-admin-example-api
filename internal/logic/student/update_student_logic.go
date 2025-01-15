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
	return &types.BaseMsgResp{Msg: data.Msg}, nil
}
