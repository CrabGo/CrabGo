// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUserToken is the golang structure for table SysUserToken.
type SysUserToken struct {
	Id                  int         `json:"id"                  orm:"Id"                  description:""` //
	CreateOn            *gtime.Time `json:"createOn"            orm:"CreateOn"            description:""` //
	Sort                int         `json:"sort"                orm:"Sort"                description:""` //
	Remark              string      `json:"remark"              orm:"Remark"              description:""` //
	UpdateOn            *gtime.Time `json:"updateOn"            orm:"UpdateOn"            description:""` //
	TenantId            string      `json:"tenantId"            orm:"TenantId"            description:""` //
	UserId              int         `json:"userId"              orm:"UserId"              description:""` //
	Token               string      `json:"token"               orm:"Token"               description:""` //
	JwtToken            string      `json:"jwtToken"            orm:"JwtToken"            description:""` //
	RefreshToken        string      `json:"refreshToken"        orm:"RefreshToken"        description:""` //
	IpAddress           string      `json:"ipAddress"           orm:"IpAddress"           description:""` //
	Country             string      `json:"country"             orm:"Country"             description:""` //
	Province            string      `json:"province"            orm:"Province"            description:""` //
	City                string      `json:"city"                orm:"City"                description:""` //
	Isp                 string      `json:"isp"                 orm:"Isp"                 description:""` //
	Browser             string      `json:"browser"             orm:"Browser"             description:""` //
	Os                  string      `json:"os"                  orm:"Os"                  description:""` //
	Device              string      `json:"device"              orm:"Device"              description:""` //
	IssuanceTime        *gtime.Time `json:"issuanceTime"        orm:"IssuanceTime"        description:""` //
	TokenExpTime        *gtime.Time `json:"tokenExpTime"        orm:"TokenExpTime"        description:""` //
	RefreshTokenExpTime *gtime.Time `json:"refreshTokenExpTime" orm:"RefreshTokenExpTime" description:""` //
	LastTime            *gtime.Time `json:"lastTime"            orm:"LastTime"            description:""` //
	UserName            string      `json:"userName"            orm:"UserName"            description:""` //
	NickName            string      `json:"nickName"            orm:"NickName"            description:""` //
	FullName            string      `json:"fullName"            orm:"FullName"            description:""` //
}
