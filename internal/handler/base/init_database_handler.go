package base

import (
	"net/http"

	"github.com/suyuan32/simple-admin-example-api/internal/logic/base"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// swagger:route get /init/database base InitDatabase
//
// Initialize database | 初始化数据库
//
// Initialize database | 初始化数据库
//
// Responses:
//  200: BaseMsgResp

func InitDatabaseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewInitDatabaseLogic(r, svcCtx)
		resp, err := l.InitDatabase()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Header.Get("Accept-Language"), err)
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
