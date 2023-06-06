package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type BookingType struct { //ประเภทของสลิป 01 = oneway 02 = roundtrip 03 = byhours
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(150)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
