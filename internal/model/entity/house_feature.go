// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HouseFeature is the golang structure for table house_feature.
type HouseFeature struct {
	Id           uint        `json:"id"           orm:"id"            description:""`         //
	HouseId      uint        `json:"houseId"      orm:"house_id"      description:"房源ID"`     // 房源ID
	FeatureType  string      `json:"featureType"  orm:"feature_type"  description:"特色类型"`     // 特色类型
	FeatureValue string      `json:"featureValue" orm:"feature_value" description:"特色值"`      // 特色值
	Weight       uint        `json:"weight"       orm:"weight"        description:"权重(1-10)"` // 权重(1-10)
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`     // 创建时间
}
