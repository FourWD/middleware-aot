package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type ExpressWayBill struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`

	UseDateTime  string  `db:"use_datetime" json:"use_datetime" gorm:"default:null; type:varchar(50); comment:'วันที่ขึ้นทางด่วน' "`
	VehicleID    string  `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	DriverID     string  `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	ExpressWayID string  `db:"expressway_id" json:"expressway_id" gorm:"type:varchar(2);"`
	SlipID       string  `db:"slip_id" json:"slip_id" gorm:"type:varchar(2);"`
	Price        float64 `db:"price" json:"price" gorm:"default:null; type:decimal; comment:'ราคาค่าทางด่วน' "`
}
