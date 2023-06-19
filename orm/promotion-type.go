package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type PromotionType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name     string `db:"name" json:"name" gorm:"type:varchar(150)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
