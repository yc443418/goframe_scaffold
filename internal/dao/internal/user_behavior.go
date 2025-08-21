// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserBehaviorDao is the data access object for the table user_behavior.
type UserBehaviorDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UserBehaviorColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UserBehaviorColumns defines and stores column names for the table user_behavior.
type UserBehaviorColumns struct {
	Id           string //
	UserId       string // 用户ID
	HouseId      string // 房源ID
	BehaviorType string // 行为类型:1-浏览,2-收藏,3-联系,4-预约看房
	Duration     string // 浏览时长(秒)
	CreatedAt    string // 行为时间
}

// userBehaviorColumns holds the columns for the table user_behavior.
var userBehaviorColumns = UserBehaviorColumns{
	Id:           "id",
	UserId:       "user_id",
	HouseId:      "house_id",
	BehaviorType: "behavior_type",
	Duration:     "duration",
	CreatedAt:    "created_at",
}

// NewUserBehaviorDao creates and returns a new DAO object for table data access.
func NewUserBehaviorDao(handlers ...gdb.ModelHandler) *UserBehaviorDao {
	return &UserBehaviorDao{
		group:    "default",
		table:    "user_behavior",
		columns:  userBehaviorColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserBehaviorDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserBehaviorDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserBehaviorDao) Columns() UserBehaviorColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserBehaviorDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserBehaviorDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserBehaviorDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
