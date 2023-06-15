package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type CompanyType struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	Type string `db:"type" json:"type" gorm:"type:varchar(150)"`
}
