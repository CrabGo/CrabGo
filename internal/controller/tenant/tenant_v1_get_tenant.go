package tenant

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"CrabGo/api/tenant/v1"
)

func (c *ControllerV1) GetTenant(ctx context.Context, req *v1.GetTenantReq) (res *v1.GetTenantRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
