package student

import (
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/student"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route post /student/list student GetStudentList
//
// Get Student list | 获取Student列表
//
// Get Student list | 获取Student列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: StudentListReq
//
// Responses:
//  200: StudentListResp

func GetStudentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StudentListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := student.NewGetStudentListLogic(r, svcCtx)
		resp, err := l.GetStudentList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
