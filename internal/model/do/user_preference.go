// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPreference is the golang structure of table user_preference for DAO operations like Where/Data.
type UserPreference struct {
	g.Meta             `orm:"table:user_preference, do:true"`
	Id                 interface{} //
	UserId             interface{} // 用户ID
	DistrictPreference interface{} // 区域偏好
	PriceMin           interface{} // 最低价格
	PriceMax           interface{} // 最高价格
	HouseType          interface{} // 户型偏好
	AreaMin            interface{} // 最小面积
	AreaMax            interface{} // 最大面积
	Tags               interface{} // 标签偏好
	Facilities         interface{} // 设施要求
	CommutingTime      interface{} // 通勤时间要求(分钟)
	CreatedAt          *gtime.Time // 创建时间
	UpdatedAt          *gtime.Time // 更新时间
}
