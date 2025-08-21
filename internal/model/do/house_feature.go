// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HouseFeature is the golang structure of table house_feature for DAO operations like Where/Data.
type HouseFeature struct {
	g.Meta       `orm:"table:house_feature, do:true"`
	Id           interface{} //
	HouseId      interface{} // 房源ID
	FeatureType  interface{} // 特色类型
	FeatureValue interface{} // 特色值
	Weight       interface{} // 权重(1-10)
	CreatedAt    *gtime.Time // 创建时间
}
