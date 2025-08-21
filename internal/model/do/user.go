// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:user, do:true"`
	Id          interface{} //
	Username    interface{} // 用户名
	Password    interface{} // 密码
	Phone       interface{} // 手机号
	Email       interface{} // 邮箱
	RealName    interface{} // 真实姓名
	Avatar      interface{} // 头像
	Gender      interface{} // 性别:0-未知,1-男,2-女
	Age         interface{} // 年龄
	Occupation  interface{} // 职业
	IncomeLevel interface{} // 收入水平:1-5级
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
