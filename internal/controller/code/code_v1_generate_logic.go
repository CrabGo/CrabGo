package code

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/code/v1"
)

func (c *ControllerV1) GenerateLogic(ctx context.Context, req *v1.GenerateLogicReq) (res *v1.GenerateLogicRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
