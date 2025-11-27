package auth

import (
	"CrabGo/internal/model"
	"CrabGo/internal/service"
	"context"

	"CrabGo/api/auth/v1"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) Auth(ctx context.Context, req *v1.AuthReq) (res *v1.AuthRes, err error) {

	var input = model.UserAuthInput{}

	err = gconv.Struct(req, &input)

	output, err := service.User().Auth(ctx, input)

	if err != nil {
		return nil, err
	}

	err = gconv.Struct(output, &res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
