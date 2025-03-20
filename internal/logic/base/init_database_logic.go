package base

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-example-api/internal/svc"
	"github.com/suyuan32/simple-admin-example-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	_, err = l.svcCtx.ExampleRpc.InitDatabase(l.ctx, &example.Empty{})
	if err != nil {
		return nil, err
	}

	err = l.insertApiData()
	if err != nil {
		return nil, err
	}

	err = l.insertMenuData()
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.Success)}, nil
}

func (l *InitDatabaseLogic) insertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:       pointy.GetPointer(uint32(1)),
		ParentId:    pointy.GetPointer(common.DefaultParentId),
		Path:        pointy.GetPointer("/example_dir"),
		Name:        pointy.GetPointer("Example"),
		Component:   pointy.GetPointer("LAYOUT"),
		Sort:        pointy.GetPointer(uint32(1)),
		Disabled:    pointy.GetPointer(false),
		ServiceName: pointy.GetPointer("Example"),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.exampleManagement"),
			Icon:               pointy.GetPointer("material-symbols:school-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:       pointy.GetPointer(uint32(2)),
		ParentId:    &menuData.Id,
		Path:        pointy.GetPointer("/example/student"),
		Name:        pointy.GetPointer("Student"),
		Component:   pointy.GetPointer("/example/student/index"),
		Sort:        pointy.GetPointer(uint32(1)),
		Disabled:    pointy.GetPointer(false),
		ServiceName: pointy.GetPointer("Example"),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.studentManagement"),
			Icon:               pointy.GetPointer("material-symbols:school-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})

	if err != nil {
		return err
	}

	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:       pointy.GetPointer(uint32(2)),
		ParentId:    &menuData.Id,
		Path:        pointy.GetPointer("/example/teacher"),
		Name:        pointy.GetPointer("Teacher"),
		Component:   pointy.GetPointer("/example/student/index"),
		Sort:        pointy.GetPointer(uint32(2)),
		Disabled:    pointy.GetPointer(false),
		ServiceName: pointy.GetPointer("Example"),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("route.teacherManagement"),
			Icon:               pointy.GetPointer("material-symbols:school-outline"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})

	if err != nil {
		return err
	}

	return nil
}
