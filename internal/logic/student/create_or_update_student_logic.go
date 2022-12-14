package student

import (
	"context"
	"net/http"

	"github.com/suyuan32/simple-admin-example-rpc/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateStudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	lang   string
}

func NewCreateOrUpdateStudentLogic(r *http.Request, svcCtx *svc.ServiceContext) *CreateOrUpdateStudentLogic {
	return &CreateOrUpdateStudentLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		svcCtx: svcCtx,
		lang:   r.Header.Get("Accept-Language"),
	}
}

func (l *CreateOrUpdateStudentLogic) CreateOrUpdateStudent(req *types.CreateOrUpdateStudentReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.CreateOrUpdateStudent(l.ctx,
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
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.lang, data.Msg)}, nil
}
