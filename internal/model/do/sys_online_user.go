// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOnlineUser is the golang structure of table SysOnlineUser for DAO operations like Where/Data.
type SysOnlineUser struct {
	g.Meta    `orm:"table:SysOnlineUser, do:true"`
	Id        any         //
	CreateOn  *gtime.Time //
	Sort      any         //
	Remark    any         //
	UpdateOn  *gtime.Time //
	SignalRId any         //
	HttpId    any         //
	UserId    any         //
	UserName  any         //
	NickName  any         //
	FullName  any         //
	Ip        any         //
	Browser   any         //
	Os        any         //
	LoginTime *gtime.Time //
	LastTime  *gtime.Time //
}
