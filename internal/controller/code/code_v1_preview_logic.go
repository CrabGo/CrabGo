package code

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/code/v1"
)

func (c *ControllerV1) PreviewLogic(ctx context.Context, req *v1.PreviewLogicReq) (res *v1.PreviewLogicRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
