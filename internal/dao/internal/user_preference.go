// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserPreferenceDao is the data access object for the table user_preference.
type UserPreferenceDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  UserPreferenceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// UserPreferenceColumns defines and stores column names for the table user_preference.
type UserPreferenceColumns struct {
	Id                 string //
	UserId             string // 用户ID
	DistrictPreference string // 区域偏好
	PriceMin           string // 最低价格
	PriceMax           string // 最高价格
	HouseType          string // 户型偏好
	AreaMin            string // 最小面积
	AreaMax            string // 最大面积
	Tags               string // 标签偏好
	Facilities         string // 设施要求
	CommutingTime      string // 通勤时间要求(分钟)
	CreatedAt          string // 创建时间
	UpdatedAt          string // 更新时间
}

// userPreferenceColumns holds the columns for the table user_preference.
var userPreferenceColumns = UserPreferenceColumns{
	Id:                 "id",
	UserId:             "user_id",
	DistrictPreference: "district_preference",
	PriceMin:           "price_min",
	PriceMax:           "price_max",
	HouseType:          "house_type",
	AreaMin:            "area_min",
	AreaMax:            "area_max",
	Tags:               "tags",
	Facilities:         "facilities",
	CommutingTime:      "commuting_time",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewUserPreferenceDao creates and returns a new DAO object for table data access.
func NewUserPreferenceDao(handlers ...gdb.ModelHandler) *UserPreferenceDao {
	return &UserPreferenceDao{
		group:    "default",
		table:    "user_preference",
		columns:  userPreferenceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserPreferenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserPreferenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserPreferenceDao) Columns() UserPreferenceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserPreferenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserPreferenceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserPreferenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
