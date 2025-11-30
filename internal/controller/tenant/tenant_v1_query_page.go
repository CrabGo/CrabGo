package tenant

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/tenant/v1"
)

func (c *ControllerV1) QueryPage(ctx context.Context, req *v1.QueryPageReq) (res *v1.QueryPageRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
