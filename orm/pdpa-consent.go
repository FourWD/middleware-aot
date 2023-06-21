package orm

import (
	"time"

	orm "github.com/FourWD/middleware/orm"
)

type PDPAConsent struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	CustomerID   string `db:"customer_id"  json:"customer_id" gorm:"type:varchar(36) ; dafault:null ;"`
	PdpaDetailID string `db:"pdpa_detail_id" json:"pdpa_detail_id" gorm:"type:varchar(36);"`

	IsAccept       bool      `db:"is_accept" json:"is_accept" gorm:"type:tinyint(2) ; comment:'ลูกค้าเป็นคนกดรับหรือไม่รับ'"`
	AcceptDatetime time.Time `db:"accept_datetime"  json:"accept_datetime" gorm:"default:null; type:Datetime; comment:'วันที่ลูกค้ากดทำรายการผ่านหน้าเว็บ/แอพ'"`

	IsCancel           bool      `db:"is_cancel" json:"is_cancel" gorm:"type:tinyint(2) ; comment:'ลูกค้าโทรมายกเลิก pdpa'"`
	PDPACancelReasonID string    `db:"pdpa_cancel_reason_id"  json:"pdpa_cancel_reason_id" gorm:"type:varchar(2) ; default:null ;"`
	CancelDatetime     time.Time `db:"cancel_datetime"  json:"cancel_datetime" gorm:"default:null; type:Datetime;"`
	Remark             string    `db:"remark"  json:"remark" gorm:"type:varchar(200) ; dafault:null ;"`
}
