// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBehavior is the golang structure for table user_behavior.
type UserBehavior struct {
	Id           uint64      `json:"id"           orm:"id"            description:""`                           //
	UserId       uint        `json:"userId"       orm:"user_id"       description:"用户ID"`                       // 用户ID
	HouseId      uint        `json:"houseId"      orm:"house_id"      description:"房源ID"`                       // 房源ID
	BehaviorType uint        `json:"behaviorType" orm:"behavior_type" description:"行为类型:1-浏览,2-收藏,3-联系,4-预约看房"` // 行为类型:1-浏览,2-收藏,3-联系,4-预约看房
	Duration     uint        `json:"duration"     orm:"duration"      description:"浏览时长(秒)"`                    // 浏览时长(秒)
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"行为时间"`                       // 行为时间
}
