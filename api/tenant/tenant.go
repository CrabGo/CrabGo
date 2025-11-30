// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package tenant

import (
	"context"

	"CrabGo/api/tenant/v1"
)

type ITenantV1 interface {
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	QueryPage(ctx context.Context, req *v1.QueryPageReq) (res *v1.QueryPageRes, err error)
	QueryList(ctx context.Context, req *v1.QueryListReq) (res *v1.QueryListRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	GetTenant(ctx context.Context, req *v1.GetTenantReq) (res *v1.GetTenantRes, err error)
}
