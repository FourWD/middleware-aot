package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Source struct { // 01 = VRP 02=BenzThonglor
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(150)"`
}
