package code

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/code/v1"
)

func (c *ControllerV1) GenerateControl(ctx context.Context, req *v1.GenerateControlReq) (res *v1.GenerateControlRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
