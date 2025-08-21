// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBehavior is the golang structure of table user_behavior for DAO operations like Where/Data.
type UserBehavior struct {
	g.Meta       `orm:"table:user_behavior, do:true"`
	Id           interface{} //
	UserId       interface{} // 用户ID
	HouseId      interface{} // 房源ID
	BehaviorType interface{} // 行为类型:1-浏览,2-收藏,3-联系,4-预约看房
	Duration     interface{} // 浏览时长(秒)
	CreatedAt    *gtime.Time // 行为时间
}
