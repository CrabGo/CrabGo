package demo

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/demo/v1"
)

func (c *ControllerV1) DeleteDemo(ctx context.Context, req *v1.DeleteDemoReq) (res *v1.DeleteDemoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
