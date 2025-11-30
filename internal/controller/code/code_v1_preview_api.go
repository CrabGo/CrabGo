package code

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/code/v1"
)

func (c *ControllerV1) PreviewApi(ctx context.Context, req *v1.PreviewApiReq) (res *v1.PreviewApiRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
