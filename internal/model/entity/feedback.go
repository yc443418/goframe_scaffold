// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Feedback is the golang structure for table feedback.
type Feedback struct {
	Id        uint        `json:"id"        orm:"id"         description:""`                           //
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户ID"`                       // 用户ID
	HouseId   uint        `json:"houseId"   orm:"house_id"   description:"房源ID"`                       // 房源ID
	Type      uint        `json:"type"      orm:"type"       description:"反馈类型:1-房源信息错误,2-推荐不准确,3-其他"` // 反馈类型:1-房源信息错误,2-推荐不准确,3-其他
	Content   string      `json:"content"   orm:"content"    description:"反馈内容"`                       // 反馈内容
	Status    uint        `json:"status"    orm:"status"     description:"状态:1-待处理,2-已处理"`             // 状态:1-待处理,2-已处理
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                       // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                       // 更新时间
}
