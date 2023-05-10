package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Address struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID string `db:"customer_id" json:"customer_id" gorm:"type:varchar(36);"`
	ProvinceID string `db:"province_id" json:"province_id" gorm:"type:varchar(2);"`
	Code       string `db:"code" json:"code" gorm:"type:varchar(20)"`
	Name       string `db:"name" json:"name" gorm:"type:varchar(150)"`
}
