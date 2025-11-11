// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLocalizerDao is the data access object for the table SysLocalizer.
type SysLocalizerDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysLocalizerColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysLocalizerColumns defines and stores column names for the table SysLocalizer.
type SysLocalizerColumns struct {
	Id       string //
	CreateOn string //
	Sort     string //
	Remark   string //
	UpdateOn string //
	Name     string //
	Value    string //
	Culture  string //
	Resource string //
}

// sysLocalizerColumns holds the columns for the table SysLocalizer.
var sysLocalizerColumns = SysLocalizerColumns{
	Id:       "Id",
	CreateOn: "CreateOn",
	Sort:     "Sort",
	Remark:   "Remark",
	UpdateOn: "UpdateOn",
	Name:     "Name",
	Value:    "Value",
	Culture:  "Culture",
	Resource: "Resource",
}

// NewSysLocalizerDao creates and returns a new DAO object for table data access.
func NewSysLocalizerDao(handlers ...gdb.ModelHandler) *SysLocalizerDao {
	return &SysLocalizerDao{
		group:    "default",
		table:    "SysLocalizer",
		columns:  sysLocalizerColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLocalizerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLocalizerDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLocalizerDao) Columns() SysLocalizerColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLocalizerDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLocalizerDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysLocalizerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
