// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Favorite is the golang structure of table favorite for DAO operations like Where/Data.
type Favorite struct {
	g.Meta    `orm:"table:favorite, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户ID
	HouseId   interface{} // 房源ID
	CreatedAt *gtime.Time // 收藏时间
}
