package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Zone struct { //no CRUD 01 = parking 02 = ลานด้านบน 03 = ตารางสนามบิน
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(50)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
