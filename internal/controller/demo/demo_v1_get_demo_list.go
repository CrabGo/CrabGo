package demo

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/demo/v1"
)

func (c *ControllerV1) GetDemoList(ctx context.Context, req *v1.GetDemoListReq) (res *v1.GetDemoListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
