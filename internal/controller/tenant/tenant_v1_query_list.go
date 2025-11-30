package tenant

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/tenant/v1"
)

func (c *ControllerV1) QueryList(ctx context.Context, req *v1.QueryListReq) (res *v1.QueryListRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
