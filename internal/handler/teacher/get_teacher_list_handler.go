package teacher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/teacher"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
)

// swagger:route post /teacher/list teacher GetTeacherList
//
// Get Teacher list | 获取Teacher列表
//
// Get Teacher list | 获取Teacher列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TeacherListReq
//
// Responses:
//  200: TeacherListResp

func GetTeacherListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeacherListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := teacher.NewGetTeacherListLogic(r, svcCtx)
		resp, err := l.GetTeacherList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
