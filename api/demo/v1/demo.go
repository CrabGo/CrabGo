package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/demo" tags:"demo" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type CreateRes struct {
	Id string `json:"id" dc:"Id"`
}

type CreateBatchReq struct {
	g.Meta `path:"/demo/batch" tags:"demo;batch" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type CreateBatchRes struct {
}

type GetListReq struct {
	g.Meta `path:"/demo/list" tags:"demo;batch" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type GetListRes struct {
}

type GetOneReq struct {
	g.Meta `path:"/demo/one" tags:"demo;batch" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type GetOneRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/demo/update" tags:"demo;batch" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type UpdateRes struct {
}

type DeleteReq struct {
	g.Meta `path:"/demo/delete" tags:"demo;batch" method:"post" summary:"通常是往数据表中插入一条据记录"`
}

type DeleteRes struct {
}
