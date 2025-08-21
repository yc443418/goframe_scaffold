// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPreference is the golang structure for table user_preference.
type UserPreference struct {
	Id                 uint        `json:"id"                 orm:"id"                  description:""`           //
	UserId             uint        `json:"userId"             orm:"user_id"             description:"用户ID"`       // 用户ID
	DistrictPreference string      `json:"districtPreference" orm:"district_preference" description:"区域偏好"`       // 区域偏好
	PriceMin           float64     `json:"priceMin"           orm:"price_min"           description:"最低价格"`       // 最低价格
	PriceMax           float64     `json:"priceMax"           orm:"price_max"           description:"最高价格"`       // 最高价格
	HouseType          string      `json:"houseType"          orm:"house_type"          description:"户型偏好"`       // 户型偏好
	AreaMin            float64     `json:"areaMin"            orm:"area_min"            description:"最小面积"`       // 最小面积
	AreaMax            float64     `json:"areaMax"            orm:"area_max"            description:"最大面积"`       // 最大面积
	Tags               string      `json:"tags"               orm:"tags"                description:"标签偏好"`       // 标签偏好
	Facilities         string      `json:"facilities"         orm:"facilities"          description:"设施要求"`       // 设施要求
	CommutingTime      uint        `json:"commutingTime"      orm:"commuting_time"      description:"通勤时间要求(分钟)"` // 通勤时间要求(分钟)
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:"创建时间"`       // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:"更新时间"`       // 更新时间
}
