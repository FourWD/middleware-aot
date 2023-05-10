package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Customer struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Code string `db:"code"  json:"code" gorm:"type:varchar(50) ; dafault:null ; index"`

	PrefixID      string `db:"prefix_id" json:"prefix_id" gorm:"type:varchar(2)"`
	Firstname     string `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname      string `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Birthday      string `db:"birthday" json:"birthday" `
	Avatar        string `db:"avatar" json:"avatar" gorm:"type:varchar(100)"`
	Email         string `db:"email" json:"email" gorm:"type:varchar(50)"`
	IDCardNo      string `db:"idcard_no" json:"idcard_no" gorm:"type:varchar(13)"`
	NationalityID string `db:"nationality_id" json:"nationality_id" gorm:"type:varchar(3)"` // table country
	PassportNo    string `db:"passport_no" json:"passport_no" gorm:"type:varchar(13)"`
	RegisterFrom  string `db:"register_from" json:"register_from" gorm:"type:varchar(2)"` // 01 = สมัครจาก counter สุวรรภูมิ 02 = จากเว็บไซต์ aot
	CompanyName   string `db:"company_name" json:"company_name" gorm:"type:varchar(150)"`
	TaxNo         string `db:"tax_no" json:"tax_no" gorm:"type:varchar(20)"`
	PhoneNo1      string `db:"phone_no1" json:"phone_no1" gorm:"type:varchar(10)"`
	PhoneNo2      string `db:"phone_no2" json:"phone_no2" gorm:"type:varchar(10)"`
}
