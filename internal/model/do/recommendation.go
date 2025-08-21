// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Recommendation is the golang structure of table recommendation for DAO operations like Where/Data.
type Recommendation struct {
	g.Meta    `orm:"table:recommendation, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户ID
	HouseId   interface{} // 房源ID
	Strategy  interface{} // 推荐策略
	Score     interface{} // 推荐分数
	Rank      interface{} // 推荐排名
	CreatedAt *gtime.Time // 推荐时间
}
