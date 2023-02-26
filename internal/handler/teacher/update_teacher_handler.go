package teacher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/teacher"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
)

// swagger:route post /teacher/update teacher UpdateTeacher
//
// Update teacher information | 更新Teacher
//
// Update teacher information | 更新Teacher
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: TeacherInfo
//
// Responses:
//  200: BaseMsgResp

func UpdateTeacherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TeacherInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := teacher.NewUpdateTeacherLogic(r, svcCtx)
		resp, err := l.UpdateTeacher(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
