// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrgDept is the golang structure of table SysOrgDept for DAO operations like Where/Data.
type SysOrgDept struct {
	g.Meta      `orm:"table:SysOrgDept, do:true"`
	Id          any         //
	CreateOn    *gtime.Time //
	Sort        any         //
	Remark      any         //
	UpdateOn    *gtime.Time //
	ParentId    any         //
	Code        any         //
	Name        any         //
	Description any         //
	CrumbCode   any         //
	CrumbDesc   any         //
	IsSystem    any         //
}
