// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysLocalizer is the golang structure for table SysLocalizer.
type SysLocalizer struct {
	Id       int         `json:"id"       orm:"Id"       description:""` //
	CreateOn *gtime.Time `json:"createOn" orm:"CreateOn" description:""` //
	Sort     int         `json:"sort"     orm:"Sort"     description:""` //
	Remark   string      `json:"remark"   orm:"Remark"   description:""` //
	UpdateOn *gtime.Time `json:"updateOn" orm:"UpdateOn" description:""` //
	Name     string      `json:"name"     orm:"Name"     description:""` //
	Value    string      `json:"value"    orm:"Value"    description:""` //
	Culture  string      `json:"culture"  orm:"Culture"  description:""` //
	Resource string      `json:"resource" orm:"Resource" description:""` //
}
