// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table SysUser.
type SysUser struct {
	Id                int         `json:"id"                orm:"Id"                description:""` //
	CreateOn          *gtime.Time `json:"createOn"          orm:"CreateOn"          description:""` //
	Sort              int         `json:"sort"              orm:"Sort"              description:""` //
	Remark            string      `json:"remark"            orm:"Remark"            description:""` //
	UpdateOn          *gtime.Time `json:"updateOn"          orm:"UpdateOn"          description:""` //
	TenantId          string      `json:"tenantId"          orm:"TenantId"          description:""` //
	UserId            string      `json:"userId"            orm:"UserId"            description:""` //
	UserName          string      `json:"userName"          orm:"UserName"          description:""` //
	NickName          string      `json:"nickName"          orm:"NickName"          description:""` //
	FullName          string      `json:"fullName"          orm:"FullName"          description:""` //
	Introduction      string      `json:"introduction"      orm:"Introduction"      description:""` //
	Signature         string      `json:"signature"         orm:"Signature"         description:""` //
	Email             string      `json:"email"             orm:"Email"             description:""` //
	EmailConfirmed    bool        `json:"emailConfirmed"    orm:"EmailConfirmed"    description:""` //
	Mobile            string      `json:"mobile"            orm:"Mobile"            description:""` //
	MobileConfirmed   bool        `json:"mobileConfirmed"   orm:"MobileConfirmed"   description:""` //
	PasswordHash      string      `json:"passwordHash"      orm:"PasswordHash"      description:""` //
	Secret            string      `json:"secret"            orm:"Secret"            description:""` //
	Avatar            string      `json:"avatar"            orm:"Avatar"            description:""` //
	IsActive          bool        `json:"isActive"          orm:"IsActive"          description:""` //
	ActiveTime        *gtime.Time `json:"activeTime"        orm:"ActiveTime"        description:""` //
	LockoutEnd        *gtime.Time `json:"lockoutEnd"        orm:"LockoutEnd"        description:""` //
	LockoutEnabled    bool        `json:"lockoutEnabled"    orm:"LockoutEnabled"    description:""` //
	AccessFailedCount int         `json:"accessFailedCount" orm:"AccessFailedCount" description:""` //
	LastLoginTime     *gtime.Time `json:"lastLoginTime"     orm:"LastLoginTime"     description:""` //
}
