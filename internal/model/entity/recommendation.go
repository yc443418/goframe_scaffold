// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Recommendation is the golang structure for table recommendation.
type Recommendation struct {
	Id        uint64      `json:"id"        orm:"id"         description:""`     //
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户ID"` // 用户ID
	HouseId   uint        `json:"houseId"   orm:"house_id"   description:"房源ID"` // 房源ID
	Strategy  string      `json:"strategy"  orm:"strategy"   description:"推荐策略"` // 推荐策略
	Score     float64     `json:"score"     orm:"score"      description:"推荐分数"` // 推荐分数
	Rank      uint        `json:"rank"      orm:"rank"       description:"推荐排名"` // 推荐排名
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"推荐时间"` // 推荐时间
}
