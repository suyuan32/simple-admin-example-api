package base

import (
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	// Student

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/student/create"),
		Description: pointy.GetPointer("apiDesc.createStudent"),
		ApiGroup:    pointy.GetPointer("student"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/student/update"),
		Description: pointy.GetPointer("apiDesc.updateStudent"),
		ApiGroup:    pointy.GetPointer("student"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/student/delete"),
		Description: pointy.GetPointer("apiDesc.deleteStudent"),
		ApiGroup:    pointy.GetPointer("student"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/student/list"),
		Description: pointy.GetPointer("apiDesc.getStudentList"),
		ApiGroup:    pointy.GetPointer("student"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/student"),
		Description: pointy.GetPointer("apiDesc.getStudentById"),
		ApiGroup:    pointy.GetPointer("student"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	// Teacher

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/teacher/create"),
		Description: pointy.GetPointer("apiDesc.createTeacher"),
		ApiGroup:    pointy.GetPointer("teacher"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/teacher/update"),
		Description: pointy.GetPointer("apiDesc.updateTeacher"),
		ApiGroup:    pointy.GetPointer("teacher"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/teacher/delete"),
		Description: pointy.GetPointer("apiDesc.deleteTeacher"),
		ApiGroup:    pointy.GetPointer("teacher"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/teacher/list"),
		Description: pointy.GetPointer("apiDesc.getTeacherList"),
		ApiGroup:    pointy.GetPointer("teacher"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Example"),
		Path:        pointy.GetPointer("/teacher"),
		Description: pointy.GetPointer("apiDesc.getTeacherById"),
		ApiGroup:    pointy.GetPointer("teacher"),
		Method:      pointy.GetPointer("POST"),
	})

	if err != nil {
		return err
	}

	return err
}
