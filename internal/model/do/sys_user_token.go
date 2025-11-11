// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserToken is the golang structure of table SysUserToken for DAO operations like Where/Data.
type SysUserToken struct {
	g.Meta              `orm:"table:SysUserToken, do:true"`
	Id                  any         //
	CreateOn            *gtime.Time //
	Sort                any         //
	Remark              any         //
	UpdateOn            *gtime.Time //
	TenantId            any         //
	UserId              any         //
	Token               any         //
	JwtToken            any         //
	RefreshToken        any         //
	IpAddress           any         //
	Country             any         //
	Province            any         //
	City                any         //
	Isp                 any         //
	Browser             any         //
	Os                  any         //
	Device              any         //
	IssuanceTime        *gtime.Time //
	TokenExpTime        *gtime.Time //
	RefreshTokenExpTime *gtime.Time //
	LastTime            *gtime.Time //
	UserName            any         //
	NickName            any         //
	FullName            any         //
}
