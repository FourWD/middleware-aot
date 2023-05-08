package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Promotion struct { //no CRUD 01 = เงินสด 02 = creditcard 03= chuqe 04 = banktransfer
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	StartDate       string `db:"start_time" json:"start_time" gorm:"type:varchar(50) ; comment:'วันที่เริ่ม' "`
	EndDate         string `db:"end_time" json:"end_time" gorm:"type:varchar(50) ; comment:'วันที่จบ'"`
	Name            string `db:"name" json:"name" gorm:"type:varchar(150)"`
	PromotionType   string `db:"promotion_type" json:"promotion_type" gorm:"type:varchar(5);"`
	PromotionAmount int    `db:"promotion_amount" json:"promotion_amount" gorm:"type:int(5);"`
}
