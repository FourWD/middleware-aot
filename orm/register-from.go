package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type RegisterFrom struct { // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Name string `db:"name" json:"name" gorm:"type:varchar(150)"`
}