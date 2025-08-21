// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Appointment is the golang structure of table appointment for DAO operations like Where/Data.
type Appointment struct {
	g.Meta          `orm:"table:appointment, do:true"`
	Id              interface{} //
	UserId          interface{} // 用户ID
	HouseId         interface{} // 房源ID
	AppointmentTime *gtime.Time // 预约时间
	ContactPhone    interface{} // 联系电话
	Status          interface{} // 状态:1-待确认,2-已确认,3-已取消,4-已完成
	Note            interface{} // 备注
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
