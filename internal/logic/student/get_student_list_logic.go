package student

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewGetStudentListLogic(r *http.Request, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *GetStudentListLogic) GetStudentList(req *types.StudentListReq) (resp *types.StudentListResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetStudentList(l.ctx,
		&example.StudentPageReq{
			Page:     req.Page,
			PageSize: req.PageSize,
			Name:     req.Name,
		})
	if err != nil {
		return nil, err
	}
	resp = &types.StudentListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.lang, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.StudentInfo{
				BaseInfo: types.BaseInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
				Name:          v.Name,
				Age:           v.Age,
				AgeInt32:      v.AgeInt32,
				AgeInt64:      v.AgeInt64,
				AgeUint:       v.AgeUint,
				AgeUint32:     v.AgeUint32,
				AgeUint64:     v.AgeUint64,
				WeightFloat:   v.WeightFloat,
				WeightFloat32: v.WeightFloat32,
				ClassId:       v.ClassId,
				EnrollAt:      v.EnrollAt,
				StatusBool:    v.StatusBool,
			})
	}
	return resp, nil
}
