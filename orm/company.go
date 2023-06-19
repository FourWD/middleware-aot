package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Company struct {
	ComapanyCode string `db:"company_code"  json:"company_code" gorm:"type:varchar(36);primary_key; "`
	orm.GormModel
	CompanyName string `db:"company_name"  json:"company_name" gorm:"type:varchar(36); "`
	ClientCode  string `db:"client_code"  json:"client_code" gorm:"type:varchar(3); "`
	ClientName  string `db:"client_name"  json:"client_name" gorm:"type:varchar(36); "`
	Address     string `db:"address" json:"address" gorm:"type:text"`
	Type        string `db:"type" json:"type" gorm:"foreignKey:CompanyRefer"`
}
