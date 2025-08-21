// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Appointment is the golang structure for table appointment.
type Appointment struct {
	Id              uint        `json:"id"              orm:"id"               description:""`                           //
	UserId          uint        `json:"userId"          orm:"user_id"          description:"用户ID"`                       // 用户ID
	HouseId         uint        `json:"houseId"         orm:"house_id"         description:"房源ID"`                       // 房源ID
	AppointmentTime *gtime.Time `json:"appointmentTime" orm:"appointment_time" description:"预约时间"`                       // 预约时间
	ContactPhone    string      `json:"contactPhone"    orm:"contact_phone"    description:"联系电话"`                       // 联系电话
	Status          uint        `json:"status"          orm:"status"           description:"状态:1-待确认,2-已确认,3-已取消,4-已完成"` // 状态:1-待确认,2-已确认,3-已取消,4-已完成
	Note            string      `json:"note"            orm:"note"             description:"备注"`                         // 备注
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"       description:"创建时间"`                       // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"       description:"更新时间"`                       // 更新时间
}
