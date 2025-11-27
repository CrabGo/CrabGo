package v1

import "github.com/gogf/gf/v2/frame/g"

type BuildInfoReq struct {
	g.Meta `path:"/build" tags:"build" method:"get" summary:"编译信息"`
}
type BuildInfoRes struct {
	Version string `json:"version"`
	GF      string `json:"gf"`
	Go      string `json:"go"`
	Git     string `json:"git"`
	Time    string `json:"time"`
}
