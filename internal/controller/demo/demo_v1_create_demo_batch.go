package demo

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/demo/v1"
)

func (c *ControllerV1) CreateDemoBatch(ctx context.Context, req *v1.CreateDemoBatchReq) (res *v1.CreateDemoBatchRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
