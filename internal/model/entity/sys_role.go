// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table SysRole.
type SysRole struct {
	Id          int         `json:"id"          orm:"Id"          description:""` //
	CreateOn    *gtime.Time `json:"createOn"    orm:"CreateOn"    description:""` //
	Sort        int         `json:"sort"        orm:"Sort"        description:""` //
	Remark      string      `json:"remark"      orm:"Remark"      description:""` //
	UpdateOn    *gtime.Time `json:"updateOn"    orm:"UpdateOn"    description:""` //
	Code        string      `json:"code"        orm:"Code"        description:""` //
	Name        string      `json:"name"        orm:"Name"        description:""` //
	Description string      `json:"description" orm:"Description" description:""` //
	DataScop    int         `json:"dataScop"    orm:"DataScop"    description:""` //
	IsSystem    bool        `json:"isSystem"    orm:"IsSystem"    description:""` //
}
