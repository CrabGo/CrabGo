// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOrgDept is the golang structure for table SysOrgDept.
type SysOrgDept struct {
	Id          int         `json:"id"          orm:"Id"          description:""` //
	CreateOn    *gtime.Time `json:"createOn"    orm:"CreateOn"    description:""` //
	Sort        int         `json:"sort"        orm:"Sort"        description:""` //
	Remark      string      `json:"remark"      orm:"Remark"      description:""` //
	UpdateOn    *gtime.Time `json:"updateOn"    orm:"UpdateOn"    description:""` //
	ParentId    int         `json:"parentId"    orm:"ParentId"    description:""` //
	Code        string      `json:"code"        orm:"Code"        description:""` //
	Name        string      `json:"name"        orm:"Name"        description:""` //
	Description string      `json:"description" orm:"Description" description:""` //
	CrumbCode   string      `json:"crumbCode"   orm:"CrumbCode"   description:""` //
	CrumbDesc   string      `json:"crumbDesc"   orm:"CrumbDesc"   description:""` //
	IsSystem    bool        `json:"isSystem"    orm:"IsSystem"    description:""` //
}
