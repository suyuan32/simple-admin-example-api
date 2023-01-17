// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "github.com/suyuan32/simple-admin-example-api/internal/handler/base"
	student "github.com/suyuan32/simple-admin-example-api/internal/handler/student"
	teacher "github.com/suyuan32/simple-admin-example-api/internal/handler/teacher"
	"github.com/suyuan32/simple-admin-example-api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/init/database",
				Handler: base.InitDatabaseHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/teacher/create_or_update",
					Handler: teacher.CreateOrUpdateTeacherHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/teacher/delete",
					Handler: teacher.DeleteTeacherHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/teacher/list",
					Handler: teacher.GetTeacherListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/teacher/batch_delete",
					Handler: teacher.BatchDeleteTeacherHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Authority},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/student/create_or_update",
					Handler: student.CreateOrUpdateStudentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/student/delete",
					Handler: student.DeleteStudentHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/student/list",
					Handler: student.GetStudentListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/student/batch_delete",
					Handler: student.BatchDeleteStudentHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
