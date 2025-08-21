// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HouseFeatureDao is the data access object for the table house_feature.
type HouseFeatureDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  HouseFeatureColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// HouseFeatureColumns defines and stores column names for the table house_feature.
type HouseFeatureColumns struct {
	Id           string //
	HouseId      string // 房源ID
	FeatureType  string // 特色类型
	FeatureValue string // 特色值
	Weight       string // 权重(1-10)
	CreatedAt    string // 创建时间
}

// houseFeatureColumns holds the columns for the table house_feature.
var houseFeatureColumns = HouseFeatureColumns{
	Id:           "id",
	HouseId:      "house_id",
	FeatureType:  "feature_type",
	FeatureValue: "feature_value",
	Weight:       "weight",
	CreatedAt:    "created_at",
}

// NewHouseFeatureDao creates and returns a new DAO object for table data access.
func NewHouseFeatureDao(handlers ...gdb.ModelHandler) *HouseFeatureDao {
	return &HouseFeatureDao{
		group:    "default",
		table:    "house_feature",
		columns:  houseFeatureColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HouseFeatureDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HouseFeatureDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HouseFeatureDao) Columns() HouseFeatureColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HouseFeatureDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HouseFeatureDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HouseFeatureDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
