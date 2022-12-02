package student

import (
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /student student CreateOrUpdateStudent
//
// Create or update Student information | 创建或更新Student
//
// Create or update Student information | 创建或更新Student
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateStudentReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateStudentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateStudentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := student.NewCreateOrUpdateStudentLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateStudent(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
