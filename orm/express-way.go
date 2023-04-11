package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type ExpressWay struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name             string  `db:"name" json:"name" gorm:"type:varchar(50)"`
	Price            float64 `db:"price" json:"price" gorm:"default:null; type:decimal; comment:'ราคาค่าทางด่วน' "`
	ExpressWayTypeID string  `db:"expressway_type_id" json:"expressway_type_id" gorm:"type:varchar(2)"`
	RowOrder         int     `db:"row_order" json:"row_order" gorm:"type:int"`
}
