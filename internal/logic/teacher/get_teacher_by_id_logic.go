package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTeacherByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIdLogic {
	return &GetTeacherByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTeacherByIdLogic) GetTeacherById(req *types.IDReq) (resp *types.TeacherInfoResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetTeacherById(l.ctx, &example.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.TeacherInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.TeacherInfo{
			BaseIDInfo: types.BaseIDInfo{
				Id:        data.Id,
				CreatedAt: data.CreatedAt,
				UpdatedAt: data.UpdatedAt,
			},
			Name: data.Name,
			Age:  data.Age,
		},
	}, nil
}
