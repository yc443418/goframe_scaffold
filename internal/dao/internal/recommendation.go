// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecommendationDao is the data access object for the table recommendation.
type RecommendationDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  RecommendationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// RecommendationColumns defines and stores column names for the table recommendation.
type RecommendationColumns struct {
	Id        string //
	UserId    string // 用户ID
	HouseId   string // 房源ID
	Strategy  string // 推荐策略
	Score     string // 推荐分数
	Rank      string // 推荐排名
	CreatedAt string // 推荐时间
}

// recommendationColumns holds the columns for the table recommendation.
var recommendationColumns = RecommendationColumns{
	Id:        "id",
	UserId:    "user_id",
	HouseId:   "house_id",
	Strategy:  "strategy",
	Score:     "score",
	Rank:      "rank",
	CreatedAt: "created_at",
}

// NewRecommendationDao creates and returns a new DAO object for table data access.
func NewRecommendationDao(handlers ...gdb.ModelHandler) *RecommendationDao {
	return &RecommendationDao{
		group:    "default",
		table:    "recommendation",
		columns:  recommendationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *RecommendationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *RecommendationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *RecommendationDao) Columns() RecommendationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *RecommendationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *RecommendationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *RecommendationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
