// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package files

import "github.com/gogf/gf/os/gtime"

type Files struct {
	Id       uint        `orm:"id,primary" json:"ID"`        // 自增ID
	CreateAt *gtime.Time `orm:"create_at"  json:"CreatedAt"` // 创建时间
	UpdateAt *gtime.Time `orm:"update_at"  json:"UpdatedAt"` // 更新时间
	DeleteAt *gtime.Time `orm:"delete_at"  json:"DeletedAt"` // 删除时间
	Name     string      `orm:"name"       json:"name"`      // 文件名
	Url      string      `orm:"url"        json:"url"`       // 文件地址
	Tag      string      `orm:"tag"        json:"tag"`       // 文件标签
	Key      string      `orm:"key"        json:"key"`       // 编号
}
