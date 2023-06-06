package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type PDPA struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Version string `db:"version" json:"version" gorm:"type:varchar(10)"`
}
