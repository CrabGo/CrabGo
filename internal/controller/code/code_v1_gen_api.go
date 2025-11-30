package code

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/code/v1"
)

func (c *ControllerV1) GenApi(ctx context.Context, req *v1.GenApiReq) (res *v1.GenApiRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
