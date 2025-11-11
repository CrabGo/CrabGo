// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOnlineUser is the golang structure for table SysOnlineUser.
type SysOnlineUser struct {
	Id        int         `json:"id"        orm:"Id"        description:""` //
	CreateOn  *gtime.Time `json:"createOn"  orm:"CreateOn"  description:""` //
	Sort      int         `json:"sort"      orm:"Sort"      description:""` //
	Remark    string      `json:"remark"    orm:"Remark"    description:""` //
	UpdateOn  *gtime.Time `json:"updateOn"  orm:"UpdateOn"  description:""` //
	SignalRId string      `json:"signalRId" orm:"SignalRId" description:""` //
	HttpId    string      `json:"httpId"    orm:"HttpId"    description:""` //
	UserId    int         `json:"userId"    orm:"UserId"    description:""` //
	UserName  string      `json:"userName"  orm:"UserName"  description:""` //
	NickName  string      `json:"nickName"  orm:"NickName"  description:""` //
	FullName  string      `json:"fullName"  orm:"FullName"  description:""` //
	Ip        string      `json:"ip"        orm:"Ip"        description:""` //
	Browser   string      `json:"browser"   orm:"Browser"   description:""` //
	Os        string      `json:"os"        orm:"Os"        description:""` //
	LoginTime *gtime.Time `json:"loginTime" orm:"LoginTime" description:""` //
	LastTime  *gtime.Time `json:"lastTime"  orm:"LastTime"  description:""` //
}
