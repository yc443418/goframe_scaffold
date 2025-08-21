// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedbackDao is the data access object for the table feedback.
type FeedbackDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FeedbackColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FeedbackColumns defines and stores column names for the table feedback.
type FeedbackColumns struct {
	Id        string //
	UserId    string // 用户ID
	HouseId   string // 房源ID
	Type      string // 反馈类型:1-房源信息错误,2-推荐不准确,3-其他
	Content   string // 反馈内容
	Status    string // 状态:1-待处理,2-已处理
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// feedbackColumns holds the columns for the table feedback.
var feedbackColumns = FeedbackColumns{
	Id:        "id",
	UserId:    "user_id",
	HouseId:   "house_id",
	Type:      "type",
	Content:   "content",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewFeedbackDao creates and returns a new DAO object for table data access.
func NewFeedbackDao(handlers ...gdb.ModelHandler) *FeedbackDao {
	return &FeedbackDao{
		group:    "default",
		table:    "feedback",
		columns:  feedbackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FeedbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FeedbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FeedbackDao) Columns() FeedbackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FeedbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FeedbackDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FeedbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
