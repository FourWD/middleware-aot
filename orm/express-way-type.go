package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type ExpressWayType struct { //no CRUD ทางด่วน 2 ประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(50)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
