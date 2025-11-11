// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserTokenDao is the data access object for the table SysUserToken.
type SysUserTokenDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysUserTokenColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysUserTokenColumns defines and stores column names for the table SysUserToken.
type SysUserTokenColumns struct {
	Id                  string //
	CreateOn            string //
	Sort                string //
	Remark              string //
	UpdateOn            string //
	TenantId            string //
	UserId              string //
	Token               string //
	JwtToken            string //
	RefreshToken        string //
	IpAddress           string //
	Country             string //
	Province            string //
	City                string //
	Isp                 string //
	Browser             string //
	Os                  string //
	Device              string //
	IssuanceTime        string //
	TokenExpTime        string //
	RefreshTokenExpTime string //
	LastTime            string //
	UserName            string //
	NickName            string //
	FullName            string //
}

// sysUserTokenColumns holds the columns for the table SysUserToken.
var sysUserTokenColumns = SysUserTokenColumns{
	Id:                  "Id",
	CreateOn:            "CreateOn",
	Sort:                "Sort",
	Remark:              "Remark",
	UpdateOn:            "UpdateOn",
	TenantId:            "TenantId",
	UserId:              "UserId",
	Token:               "Token",
	JwtToken:            "JwtToken",
	RefreshToken:        "RefreshToken",
	IpAddress:           "IpAddress",
	Country:             "Country",
	Province:            "Province",
	City:                "City",
	Isp:                 "Isp",
	Browser:             "Browser",
	Os:                  "Os",
	Device:              "Device",
	IssuanceTime:        "IssuanceTime",
	TokenExpTime:        "TokenExpTime",
	RefreshTokenExpTime: "RefreshTokenExpTime",
	LastTime:            "LastTime",
	UserName:            "UserName",
	NickName:            "NickName",
	FullName:            "FullName",
}

// NewSysUserTokenDao creates and returns a new DAO object for table data access.
func NewSysUserTokenDao(handlers ...gdb.ModelHandler) *SysUserTokenDao {
	return &SysUserTokenDao{
		group:    "default",
		table:    "SysUserToken",
		columns:  sysUserTokenColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserTokenDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserTokenDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserTokenDao) Columns() SysUserTokenColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserTokenDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserTokenDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysUserTokenDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
