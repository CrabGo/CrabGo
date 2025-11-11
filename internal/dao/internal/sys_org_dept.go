// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrgDeptDao is the data access object for the table SysOrgDept.
type SysOrgDeptDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysOrgDeptColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysOrgDeptColumns defines and stores column names for the table SysOrgDept.
type SysOrgDeptColumns struct {
	Id          string //
	CreateOn    string //
	Sort        string //
	Remark      string //
	UpdateOn    string //
	ParentId    string //
	Code        string //
	Name        string //
	Description string //
	CrumbCode   string //
	CrumbDesc   string //
	IsSystem    string //
}

// sysOrgDeptColumns holds the columns for the table SysOrgDept.
var sysOrgDeptColumns = SysOrgDeptColumns{
	Id:          "Id",
	CreateOn:    "CreateOn",
	Sort:        "Sort",
	Remark:      "Remark",
	UpdateOn:    "UpdateOn",
	ParentId:    "ParentId",
	Code:        "Code",
	Name:        "Name",
	Description: "Description",
	CrumbCode:   "CrumbCode",
	CrumbDesc:   "CrumbDesc",
	IsSystem:    "IsSystem",
}

// NewSysOrgDeptDao creates and returns a new DAO object for table data access.
func NewSysOrgDeptDao(handlers ...gdb.ModelHandler) *SysOrgDeptDao {
	return &SysOrgDeptDao{
		group:    "default",
		table:    "SysOrgDept",
		columns:  sysOrgDeptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOrgDeptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOrgDeptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOrgDeptDao) Columns() SysOrgDeptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOrgDeptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOrgDeptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysOrgDeptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
