// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppointmentDao is the data access object for the table appointment.
type AppointmentDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppointmentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppointmentColumns defines and stores column names for the table appointment.
type AppointmentColumns struct {
	Id              string //
	UserId          string // 用户ID
	HouseId         string // 房源ID
	AppointmentTime string // 预约时间
	ContactPhone    string // 联系电话
	Status          string // 状态:1-待确认,2-已确认,3-已取消,4-已完成
	Note            string // 备注
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// appointmentColumns holds the columns for the table appointment.
var appointmentColumns = AppointmentColumns{
	Id:              "id",
	UserId:          "user_id",
	HouseId:         "house_id",
	AppointmentTime: "appointment_time",
	ContactPhone:    "contact_phone",
	Status:          "status",
	Note:            "note",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewAppointmentDao creates and returns a new DAO object for table data access.
func NewAppointmentDao(handlers ...gdb.ModelHandler) *AppointmentDao {
	return &AppointmentDao{
		group:    "default",
		table:    "appointment",
		columns:  appointmentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppointmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppointmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppointmentDao) Columns() AppointmentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppointmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppointmentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppointmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
