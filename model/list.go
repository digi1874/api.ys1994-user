/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-04 17:45:55
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-04-03 18:29:00
 */

package model

import (
	"github.com/jinzhu/gorm"
)

// Filters 过滤
type Filters struct {
	IDs              []uint   `json:"ids"`
	CreatedTimeStart uint     `json:"createdTimeStart"`
	CreatedTimeEnd   uint     `json:"createdTimeEnd"`
	UpdatedTimeStart uint     `json:"updatedTimeStart"`
	UpdatedTimeEnd   uint     `json:"updatedTimeEnd"`
}

// List 列表
type List struct {
	Count           int         `json:"count"`
	Page            int         `json:"page"`
	Size            int         `json:"size"`
	Filters         Filters     `json:"-"`
	Orders          [][2]string `json:"-"`
}

func (f Filters) dbHandle(db *gorm.DB, tableName string) *gorm.DB {
	if len(f.IDs) > 0 {
		db = db.Where("`" + tableName + "`.`id` IN (?)", f.IDs)
	}
	if f.CreatedTimeStart > 0 {
		db = db.Where("`" + tableName + "`.`created_time` >= ?", f.CreatedTimeStart)
	}
	if f.CreatedTimeEnd > 0 {
		db = db.Where("`" + tableName + "`.`created_time` <= ?", f.CreatedTimeEnd)
	}
	if f.UpdatedTimeStart > 0 {
		db = db.Where("`" + tableName + "`.`updated_time` >= ?", f.UpdatedTimeStart)
	}
	if f.UpdatedTimeEnd > 0 {
		db = db.Where("`" + tableName + "`.`updated_time` <= ?", f.UpdatedTimeEnd)
	}
	return db
}

func (l *List) dbHandle(db *gorm.DB) *gorm.DB {
	for _, v := range l.Orders {
		switch v[1] {
			case "ASC": db = db.Order("`" + v[0] + "` ASC")
			case "DESC": db = db.Order("`" + v[0] + "` DESC")
		}
	}

	if l.Page < 1 {
		l.Page = 1
	}
	if l.Size < 1 {
		l.Size = 20
	}
	offset := (l.Page - 1) * l.Size
	return db.Limit(l.Size).Offset(offset)
}
