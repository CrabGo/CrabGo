package v1

import "github.com/gogf/gf/v2/frame/g"

type QueryDatabaseReq struct {
	g.Meta `path:"/code/database" tags:"code" method:"get" summary:"查询数据库"`
}

type QueryDatabaseRes []string

type QueryTablesReq struct {
	g.Meta   `path:"/code/{database}/tables" tags:"code" method:"get" summary:"查询数据库表"`
	database string
}

type QueryTablesRes struct {
}

type QueryColumnsReq struct {
	g.Meta   `path:"/code/{database}/{table}/columns" tags:"code" method:"get" summary:"查询数据库表字段"`
	database string
	table    string
}

type PreviewApiReq struct {
	g.Meta   `path:"/code/{database}/{table}/preview/api" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type PreviewApiRes struct {
}

type PreviewControlReq struct {
	g.Meta   `path:"/code/{database}/{table}/preview/control" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type PreviewControlRes struct {
}
type PreviewLogicReq struct {
	g.Meta   `path:"/code/{database}/{table}/preview/logic" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type PreviewLogicRes struct {
}

type GenerateApiReq struct {
	g.Meta   `path:"/code/{database}/{table}/generate/api" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type GenerateApiRes struct{}

type GenerateControlReq struct {
	g.Meta   `path:"/code/{database}/{table}/generate/control" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type GenerateControlRes struct{}

type GenerateLogicReq struct {
	g.Meta   `path:"/code/{database}/{table}/generate/logic" tags:"code" method:"post" summary:"预览API代码"`
	database string
	table    string
}

type GenerateLogicRes struct{}
