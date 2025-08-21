// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HouseDao is the data access object for the table house.
type HouseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HouseColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HouseColumns defines and stores column names for the table house.
type HouseColumns struct {
	Id           string //
	Title        string // 房源标题
	Description  string // 房源描述
	Price        string // 月租金
	Area         string // 面积(㎡)
	HouseType    string // 户型
	Floor        string // 楼层
	TotalFloors  string // 总楼层
	District     string // 区域
	Address      string // 详细地址
	Longitude    string // 经度
	Latitude     string // 纬度
	Facilities   string // 设施
	Tags         string // 标签
	MainImage    string // 主图
	Images       string // 图片列表
	ContactName  string // 联系人
	ContactPhone string // 联系电话
	Status       string // 状态:1-待租,2-已租,3-下架
	ViewCount    string // 浏览数
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// houseColumns holds the columns for the table house.
var houseColumns = HouseColumns{
	Id:           "id",
	Title:        "title",
	Description:  "description",
	Price:        "price",
	Area:         "area",
	HouseType:    "house_type",
	Floor:        "floor",
	TotalFloors:  "total_floors",
	District:     "district",
	Address:      "address",
	Longitude:    "longitude",
	Latitude:     "latitude",
	Facilities:   "facilities",
	Tags:         "tags",
	MainImage:    "main_image",
	Images:       "images",
	ContactName:  "contact_name",
	ContactPhone: "contact_phone",
	Status:       "status",
	ViewCount:    "view_count",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewHouseDao creates and returns a new DAO object for table data access.
func NewHouseDao(handlers ...gdb.ModelHandler) *HouseDao {
	return &HouseDao{
		group:    "default",
		table:    "house",
		columns:  houseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HouseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HouseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HouseDao) Columns() HouseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HouseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HouseDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HouseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
