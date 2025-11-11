// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLocalizer is the golang structure of table SysLocalizer for DAO operations like Where/Data.
type SysLocalizer struct {
	g.Meta   `orm:"table:SysLocalizer, do:true"`
	Id       any         //
	CreateOn *gtime.Time //
	Sort     any         //
	Remark   any         //
	UpdateOn *gtime.Time //
	Name     any         //
	Value    any         //
	Culture  any         //
	Resource any         //
}
