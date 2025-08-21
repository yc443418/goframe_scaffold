// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// House is the golang structure of table house for DAO operations like Where/Data.
type House struct {
	g.Meta       `orm:"table:house, do:true"`
	Id           interface{} //
	Title        interface{} // 房源标题
	Description  interface{} // 房源描述
	Price        interface{} // 月租金
	Area         interface{} // 面积(㎡)
	HouseType    interface{} // 户型
	Floor        interface{} // 楼层
	TotalFloors  interface{} // 总楼层
	District     interface{} // 区域
	Address      interface{} // 详细地址
	Longitude    interface{} // 经度
	Latitude     interface{} // 纬度
	Facilities   interface{} // 设施
	Tags         interface{} // 标签
	MainImage    interface{} // 主图
	Images       interface{} // 图片列表
	ContactName  interface{} // 联系人
	ContactPhone interface{} // 联系电话
	Status       interface{} // 状态:1-待租,2-已租,3-下架
	ViewCount    interface{} // 浏览数
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
