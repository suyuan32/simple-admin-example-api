package teacher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/teacher"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"
)

// swagger:route post /teacher/create_or_update teacher CreateOrUpdateTeacher
//
// Create or update Teacher information | 创建或更新Teacher
//
// Create or update Teacher information | 创建或更新Teacher
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: CreateOrUpdateTeacherReq
//
// Responses:
//  200: BaseMsgResp

func CreateOrUpdateTeacherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateOrUpdateTeacherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := teacher.NewCreateOrUpdateTeacherLogic(r, svcCtx)
		resp, err := l.CreateOrUpdateTeacher(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
