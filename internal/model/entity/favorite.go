// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Favorite is the golang structure for table favorite.
type Favorite struct {
	Id        uint        `json:"id"        orm:"id"         description:""`     //
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户ID"` // 用户ID
	HouseId   uint        `json:"houseId"   orm:"house_id"   description:"房源ID"` // 房源ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"收藏时间"` // 收藏时间
}
