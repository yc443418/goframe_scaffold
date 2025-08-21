// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          uint        `json:"id"          orm:"id"           description:""`                //
	Username    string      `json:"username"    orm:"username"     description:"用户名"`             // 用户名
	Password    string      `json:"password"    orm:"password"     description:"密码"`              // 密码
	Phone       string      `json:"phone"       orm:"phone"        description:"手机号"`             // 手机号
	Email       string      `json:"email"       orm:"email"        description:"邮箱"`              // 邮箱
	RealName    string      `json:"realName"    orm:"real_name"    description:"真实姓名"`            // 真实姓名
	Avatar      string      `json:"avatar"      orm:"avatar"       description:"头像"`              // 头像
	Gender      int         `json:"gender"      orm:"gender"       description:"性别:0-未知,1-男,2-女"` // 性别:0-未知,1-男,2-女
	Age         uint        `json:"age"         orm:"age"          description:"年龄"`              // 年龄
	Occupation  string      `json:"occupation"  orm:"occupation"   description:"职业"`              // 职业
	IncomeLevel uint        `json:"incomeLevel" orm:"income_level" description:"收入水平:1-5级"`       // 收入水平:1-5级
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`            // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`            // 更新时间
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`            // 删除时间
}
