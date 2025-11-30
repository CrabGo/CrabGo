package v1

import (
	"CrabGo/api"
	"CrabGo/internal/consts"
	"CrabGo/internal/consts/product"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CreateReq 创建租户
type CreateReq struct {
	g.Meta                  `path:"/tenant" tags:"tenant" method:"post" summary:"创建租户"`
	Code                    string          `json:"code" v:"required#请输入租户便秘"  dc:"租户编码"`
	Name                    string          `json:"name" v:"required#请输入租户名称" dc:"租户名称"`
	Version                 product.Version `json:"version"  v:"required#请输入产品版本" dc:"产品版本"`
	Status                  consts.Status   `json:"status"  v:"bail|required|enums" dc:"租户状态"`
	DbProvider              consts.Provider `json:"dbProvider" v:"bail|required|enums" dc:"配置数据库类型"`
	DbConnectionString      string          `json:"dbConnectionString" dc:"配置数据库连接字符串"`
	CacheProvider           consts.Provider `json:"cacheProvider"  c:"缓存配置类型"`
	CacheConnectionString   string          `json:"cacheConnectionString" dc:"缓存配置连接字符串"`
	HistoryProvider         consts.Provider `json:"historyProvider" dc:"历史数据库类型"`
	HistoryConnectionString string          `json:"historyConnectionString"  dc:"历史数据库连接字符串"`
	ContactName             string          `json:"contactName" dc:"联系人"`
	Email                   string          `json:"email" v:"email"  dc:"邮件地址"`
	Phone                   string          `json:"phone"  v:"phone" dc:"联系电话"`
	Address                 string          `json:"address" dc:"联系地址"`
	Site                    string          `json:"site"   v:"domain"   dc:"网站"`
	Callback                string          `json:"callback"    dc:"回调地址"`
	Remark                  string          `json:"remark"     dc:"备注"`
	Sort                    int             `json:"sort"        d:"100"   dc:"排序"`
}

// CreateRes 创建租户
type CreateRes struct {
	TenantId   string `json:"tenantId" dc:"租户ID"`
	Code       string `json:"code"   dc:"租户编码"`
	Name       string `json:"name"   dc:"租户名称"`
	Version    string `json:"version" dc:"产品版本"`
	Status     int    `json:"status"  dc:"租户状态"`
	StatusDesc string `json:"status_desc" dc:"租户状态描述"`
}

type QueryPageReq struct {
	g.Meta `path:"/tenant/page" tags:"tenant" method:"post" summary:"查询租户"`
	api.PaginationReq
	Code        string          `json:"code"  dc:"租户编码"`
	Name        string          `json:"name" dc:"租户名称"`
	Version     product.Version `json:"version"   dc:"产品版本"`
	Status      consts.Status   `json:"status"  dc:"租户状态"`
	ContactName string          `json:"contactName" dc:"联系人"`
	Email       string          `json:"email"   dc:"邮件地址"`
	Phone       string          `json:"phone"   dc:"联系电话"`
	Address     string          `json:"address" dc:"联系地址"`
	Site        string          `json:"site"     dc:"网站"`
	Callback    string          `json:"callback"    dc:"回调地址"`
	Remark      string          `json:"remark"     dc:"备注"`
}

type TenantViewModel struct {
	Id                      int             `json:"id"`
	CreateOn                *gtime.Time     `json:"createOn"`
	Sort                    int             `json:"sort"`
	Remark                  string          `json:"remark"`
	UpdateOn                *gtime.Time     `json:"updateOn"`
	TenantId                string          `json:"tenantId"`
	Code                    string          `json:"code" `
	Name                    string          `json:"name"`
	Version                 string          `json:"version"`
	Status                  consts.Status   `json:"status"`
	StatusDesc              string          `json:"status_desc"`
	DbProvider              consts.Provider `json:"dbProvider"`
	DbConnectionString      string          `json:"dbConnectionString"`
	CacheProvider           consts.Provider `json:"cacheProvider"`
	CacheConnectionString   string          `json:"cacheConnectionString"`
	HistoryProvider         consts.Provider `json:"historyProvider"`
	HistoryConnectionString string          `json:"historyConnectionString"`
	CreateTime              *gtime.Time     `json:"createTime"`
	ContactName             string          `json:"contactName"`
	Email                   string          `json:"email"`
	Phone                   string          `json:"phone"`
	Address                 string          `json:"address"`
	Site                    string          `json:"site"`
	Callback                string          `json:"callback"`
}

type QueryPageRes api.PaginationRes

type QueryListReq struct {
	g.Meta `path:"/tenant/page" tags:"tenant" method:"post" summary:"查询租户"`
	api.PaginationReq
	Code        string          `json:"code"  dc:"租户编码"`
	Name        string          `json:"name" dc:"租户名称"`
	Version     product.Version `json:"version"   dc:"产品版本"`
	Status      consts.Status   `json:"status"  dc:"租户状态"`
	ContactName string          `json:"contactName" dc:"联系人"`
	Email       string          `json:"email"   dc:"邮件地址"`
	Phone       string          `json:"phone"   dc:"联系电话"`
	Address     string          `json:"address" dc:"联系地址"`
	Site        string          `json:"site"     dc:"网站"`
	Callback    string          `json:"callback"    dc:"回调地址"`
	Remark      string          `json:"remark"     dc:"备注"`
}

type QueryListRes []TenantViewModel

type DeleteReq struct {
	g.Meta `path:"/tenant/{id}" tags:"tenant" method:"delete" summary:"删除租户"`
	Id     int `json:"id"`
}

type DeleteRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/tenant/{id}" tags:"tenant" method:"post" summary:"删除租户"`
	Id     int `json:"id"`
	TenantViewModel
}

type UpdateRes struct{}

type GetTenantReq struct {
	g.Meta `path:"/tenant/{id}" tags:"tenant" method:"get" summary:"查询租户"`
	Id     int `json:"id"`
}

type GetTenantRes TenantViewModel
