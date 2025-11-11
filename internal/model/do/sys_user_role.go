// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure of table SysUserRole for DAO operations like Where/Data.
type SysUserRole struct {
	g.Meta   `orm:"table:SysUserRole, do:true"`
	Id       any         //
	CreateOn *gtime.Time //
	Sort     any         //
	Remark   any         //
	UpdateOn *gtime.Time //
	UserId   any         //
	RoleId   any         //
}
