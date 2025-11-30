// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package code

import (
	"context"

	"CrabGo/api/code/v1"
)

type ICodeV1 interface {
	QueryDatabase(ctx context.Context, req *v1.QueryDatabaseReq) (res *v1.QueryDatabaseRes, err error)
	QueryTables(ctx context.Context, req *v1.QueryTablesReq) (res *v1.QueryTablesRes, err error)
	QueryColumns(ctx context.Context, req *v1.QueryColumnsReq) (res *v1.QueryColumnsRes, err error)
	PreviewApi(ctx context.Context, req *v1.PreviewApiReq) (res *v1.PreviewApiRes, err error)
	PreviewControl(ctx context.Context, req *v1.PreviewControlReq) (res *v1.PreviewControlRes, err error)
	PreviewLogic(ctx context.Context, req *v1.PreviewLogicReq) (res *v1.PreviewLogicRes, err error)
	GenerateApi(ctx context.Context, req *v1.GenerateApiReq) (res *v1.GenerateApiRes, err error)
	GenerateControl(ctx context.Context, req *v1.GenerateControlReq) (res *v1.GenerateControlRes, err error)
	GenerateLogic(ctx context.Context, req *v1.GenerateLogicReq) (res *v1.GenerateLogicRes, err error)
}
