// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// House is the golang structure for table house.
type House struct {
	Id           uint        `json:"id"           orm:"id"            description:""`                  //
	Title        string      `json:"title"        orm:"title"         description:"房源标题"`              // 房源标题
	Description  string      `json:"description"  orm:"description"   description:"房源描述"`              // 房源描述
	Price        float64     `json:"price"        orm:"price"         description:"月租金"`               // 月租金
	Area         float64     `json:"area"         orm:"area"          description:"面积(㎡)"`             // 面积(㎡)
	HouseType    string      `json:"houseType"    orm:"house_type"    description:"户型"`                // 户型
	Floor        string      `json:"floor"        orm:"floor"         description:"楼层"`                // 楼层
	TotalFloors  uint        `json:"totalFloors"  orm:"total_floors"  description:"总楼层"`               // 总楼层
	District     string      `json:"district"     orm:"district"      description:"区域"`                // 区域
	Address      string      `json:"address"      orm:"address"       description:"详细地址"`              // 详细地址
	Longitude    float64     `json:"longitude"    orm:"longitude"     description:"经度"`                // 经度
	Latitude     float64     `json:"latitude"     orm:"latitude"      description:"纬度"`                // 纬度
	Facilities   string      `json:"facilities"   orm:"facilities"    description:"设施"`                // 设施
	Tags         string      `json:"tags"         orm:"tags"          description:"标签"`                // 标签
	MainImage    string      `json:"mainImage"    orm:"main_image"    description:"主图"`                // 主图
	Images       string      `json:"images"       orm:"images"        description:"图片列表"`              // 图片列表
	ContactName  string      `json:"contactName"  orm:"contact_name"  description:"联系人"`               // 联系人
	ContactPhone string      `json:"contactPhone" orm:"contact_phone" description:"联系电话"`              // 联系电话
	Status       uint        `json:"status"       orm:"status"        description:"状态:1-待租,2-已租,3-下架"` // 状态:1-待租,2-已租,3-下架
	ViewCount    uint        `json:"viewCount"    orm:"view_count"    description:"浏览数"`               // 浏览数
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`              // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`              // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`              // 删除时间
}
