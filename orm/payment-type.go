package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type PaymentType struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name            string `db:"name" json:"name" gorm:"type:varchar(150)"`
	IsWht           bool   `db:"is_wht" json:"is_wht" gorm:"default:0; type:tinyint(1); comment:'มีwht?' "`
	PromotionType   string `db:"promotion_type" json:"promotion_type" gorm:"type:varchar(5);"`
	PromotionAmount int    `db:"promotion_amount" json:"promotion_amount" gorm:"type:int(5);"`
	RowOrder        int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
