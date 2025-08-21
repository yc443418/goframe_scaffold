// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Feedback is the golang structure of table feedback for DAO operations like Where/Data.
type Feedback struct {
	g.Meta    `orm:"table:feedback, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户ID
	HouseId   interface{} // 房源ID
	Type      interface{} // 反馈类型:1-房源信息错误,2-推荐不准确,3-其他
	Content   interface{} // 反馈内容
	Status    interface{} // 状态:1-待处理,2-已处理
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
