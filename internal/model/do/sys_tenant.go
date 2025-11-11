// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenant is the golang structure of table SysTenant for DAO operations like Where/Data.
type SysTenant struct {
	g.Meta                  `orm:"table:SysTenant, do:true"`
	Id                      any         //
	CreateOn                *gtime.Time //
	Sort                    any         //
	Remark                  any         //
	UpdateOn                *gtime.Time //
	TenantId                any         //
	Code                    any         //
	Name                    any         //
	Version                 any         //
	Status                  any         //
	DbProvider              any         //
	DbConnectionString      any         //
	CacheProvider           any         //
	CacheConnectionString   any         //
	HistoryProvider         any         //
	HistoryConnectionString any         //
	CreateTime              *gtime.Time //
	ContactName             any         //
	Email                   any         //
	Phone                   any         //
	Address                 any         //
	Site                    any         //
	Callback                any         //
}
