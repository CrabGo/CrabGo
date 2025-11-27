package user

import (
	"CrabGo/internal/dao"
	"CrabGo/internal/model"
	"CrabGo/internal/model/entity"
	"CrabGo/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sUser struct{}

func New() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(New())
}

func (s *sUser) Auth(ctx context.Context, input model.UserAuthInput) (out *model.UserAuthOutput, err error) {

	return nil, nil
}

func (s *sUser) GetById(ctx context.Context, userId string) (user *entity.SysUser, err error) {

	count, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Count()
	if err != nil {
		return nil, gerror.New("数据库查询异常,请稍后重试")
	}

	if count <= 0 {
		return nil, gerror.Newf("用户%v不存在", userId)
	}

	err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().UserId, userId).Scan(&user)

	if err != nil {
		return nil, gerror.New("查询用户信息失败")
	}

	return user, nil
}
