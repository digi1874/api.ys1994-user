/*
 * @Author: lin.zhenhui
 * @Date: 2020-03-20 20:53:17
 * @Last Modified by: lin.zhenhui
 * @Last Modified time: 2020-05-13 19:45:09
 */

package model

import (
	"github.com/jinzhu/gorm"

	"api.ys1994-user/database"
)

type vod struct {
	ID              uint               `json:"id"`
	Name            string             `json:"name"`
}

// Vod Vod
type Vod struct {
	Base
	ID              uint               `json:"id"`
	Actor           string             `json:"actor"`
	Area            string             `json:"area"`
	Serial          string             `json:"serial"`
	Director        string             `json:"director"`
	Name            string             `json:"name"`
	Pic             string             `json:"pic"`
	Year            uint16             `json:"year"`
	TypeID          uint8              `json:"typeId"`
	Filter          database.Vod       `json:"-"`
}

type vodList struct {
	List
	Data         []vod          `json:"data"`
	Filter       database.Vod   `json:"-"`
}

// Super Super
func (v *Vod) Super() {
	v.Base.Super(v, &v.Filter)
}

func handleFilterVod(db *gorm.DB, filter database.Vod) *gorm.DB {
	if filter.Name != "" {
		name := "%" + filter.Name + "%"
		db = db.Where("`vods`.`name` LIKE ? OR `vods`.`sub_name` LIKE ?", name, name)
		filter.Name = ""
	}

	if filter.Actor != "" {
		db = db.Where("`vods`.`actor` LIKE ?", "%" + filter.Actor + "%")
		filter.Actor = ""
	}

	if filter.Director != "" {
		db = db.Where("`vods`.`director` LIKE ?", "%" + filter.Director + "%")
		filter.Director = ""
	}

	if filter.Serial != "" {
		db = db.Where("`vods`.`serial` LIKE ?", "%" + filter.Serial + "%")
		filter.Serial = ""
	}

	if filter.Area != "" {
		db = db.Where("`vods`.`area` LIKE ?", "%" + filter.Area + "%")
		filter.Area = ""
	}

	if filter.Lang != "" {
		db = db.Where("`vods`.`lang` LIKE ?", "%" + filter.Lang + "%")
		filter.Lang = ""
	}

	return db.Where(&filter)
}

// Find 列表
func (vl *vodList) Find() {
	db := vl.Filters.dbHandle(database.DB, "vods")
	vl.dbHandle(db.Where(&vl.Filter)).Find(&vl.Data)
}
