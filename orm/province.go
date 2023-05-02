package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Province struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code string `db:"code" json:"code" gorm:"type:varchar(20)"`
	Name string `db:"name" json:"name" gorm:"type:varchar(150)"`
}
