// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure for table SysTenant.
type SysTenant struct {
	Id                      int         `json:"id"                      orm:"Id"                      description:""` //
	CreateOn                *gtime.Time `json:"createOn"                orm:"CreateOn"                description:""` //
	Sort                    int         `json:"sort"                    orm:"Sort"                    description:""` //
	Remark                  string      `json:"remark"                  orm:"Remark"                  description:""` //
	UpdateOn                *gtime.Time `json:"updateOn"                orm:"UpdateOn"                description:""` //
	TenantId                string      `json:"tenantId"                orm:"TenantId"                description:""` //
	Code                    string      `json:"code"                    orm:"Code"                    description:""` //
	Name                    string      `json:"name"                    orm:"Name"                    description:""` //
	Version                 string      `json:"version"                 orm:"Version"                 description:""` //
	Status                  int         `json:"status"                  orm:"Status"                  description:""` //
	DbProvider              string      `json:"dbProvider"              orm:"DbProvider"              description:""` //
	DbConnectionString      string      `json:"dbConnectionString"      orm:"DbConnectionString"      description:""` //
	CacheProvider           string      `json:"cacheProvider"           orm:"CacheProvider"           description:""` //
	CacheConnectionString   string      `json:"cacheConnectionString"   orm:"CacheConnectionString"   description:""` //
	HistoryProvider         string      `json:"historyProvider"         orm:"HistoryProvider"         description:""` //
	HistoryConnectionString string      `json:"historyConnectionString" orm:"HistoryConnectionString" description:""` //
	CreateTime              *gtime.Time `json:"createTime"              orm:"CreateTime"              description:""` //
	ContactName             string      `json:"contactName"             orm:"ContactName"             description:""` //
	Email                   string      `json:"email"                   orm:"Email"                   description:""` //
	Phone                   string      `json:"phone"                   orm:"Phone"                   description:""` //
	Address                 string      `json:"address"                 orm:"Address"                 description:""` //
	Site                    string      `json:"site"                    orm:"Site"                    description:""` //
	Callback                string      `json:"callback"                orm:"Callback"                description:""` //
}
