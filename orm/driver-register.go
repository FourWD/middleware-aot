package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type DriverRegister struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	DriverID          string `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	VehicleID         string `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	IsGetKey          bool   `db:"is_getkey" json:"is_getkey" gorm:"default:0; type:tinyint(1); comment:'รับกุญแจยัง' "`
	GetKeyDateTime    string `db:"getkey_datetime"  json:"getkey_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลารับกุญแจ' "`
	GetKeyApprover    string `db:"getkey_approver"  json:"getkey_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติรับกุญแจ' "`
	IsReturnKey       bool   `db:"is_returnkey"  json:"is_returnkey" gorm:"default:0; type:tinyint(1); comment:'คืนกุญแจยัง' "`
	ReturnKeyDateTime string `db:"returnkey_datetime"  json:"returnkey_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาคืนกุญแจ' "`
	ReturnKeyApprover string `db:"returnkey_approver"  json:"returnkey_approver" gorm:"default:null; type:varchar(50); comment:'ผู้อนุมัติคืนกุญแจ' "`
	WorkShiftID       string `db:"workshift_id"  json:"workshift_id" gorm:"default:null; type:varchar(2); "`
}
