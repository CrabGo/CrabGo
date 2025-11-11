// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table SysUser for DAO operations like Where/Data.
type SysUser struct {
	g.Meta            `orm:"table:SysUser, do:true"`
	Id                any         //
	CreateOn          *gtime.Time //
	Sort              any         //
	Remark            any         //
	UpdateOn          *gtime.Time //
	TenantId          any         //
	UserId            any         //
	UserName          any         //
	NickName          any         //
	FullName          any         //
	Introduction      any         //
	Signature         any         //
	Email             any         //
	EmailConfirmed    any         //
	Mobile            any         //
	MobileConfirmed   any         //
	PasswordHash      any         //
	Secret            any         //
	Avatar            any         //
	IsActive          any         //
	ActiveTime        *gtime.Time //
	LockoutEnd        *gtime.Time //
	LockoutEnabled    any         //
	AccessFailedCount any         //
	LastLoginTime     *gtime.Time //
}
