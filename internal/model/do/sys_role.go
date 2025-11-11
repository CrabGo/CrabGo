// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure of table SysRole for DAO operations like Where/Data.
type SysRole struct {
	g.Meta      `orm:"table:SysRole, do:true"`
	Id          any         //
	CreateOn    *gtime.Time //
	Sort        any         //
	Remark      any         //
	UpdateOn    *gtime.Time //
	Code        any         //
	Name        any         //
	Description any         //
	DataScop    any         //
	IsSystem    any         //
}
