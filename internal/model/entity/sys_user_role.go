// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserRole is the golang structure for table SysUserRole.
type SysUserRole struct {
	Id       int         `json:"id"       orm:"Id"       description:""` //
	CreateOn *gtime.Time `json:"createOn" orm:"CreateOn" description:""` //
	Sort     int         `json:"sort"     orm:"Sort"     description:""` //
	Remark   string      `json:"remark"   orm:"Remark"   description:""` //
	UpdateOn *gtime.Time `json:"updateOn" orm:"UpdateOn" description:""` //
	UserId   string      `json:"userId"   orm:"UserId"   description:""` //
	RoleId   string      `json:"roleId"   orm:"RoleId"   description:""` //
}
