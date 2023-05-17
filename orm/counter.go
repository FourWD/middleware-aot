package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Counter struct { // จุดขายหน้าเค้าเตอร์  ที่สุวรรณภูมิ
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Name   string `db:"name" json:"name" gorm:"type:varchar(150)"`
	IP     string `db:"ip" json:"ip" gorm:"type:varchar(20)"`
	Remark string `db:"remark" json:"remark" gorm:"type:varchar(200)"`
}
