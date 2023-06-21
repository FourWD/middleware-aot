package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type ExpresswayBill struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`

	UseAt        string  `db:"use_at" json:"use_at" gorm:"default:null; type:varchar(50); comment:'วันที่ขึ้นทางด่วน' "`
	VehicleID    string  `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	DriverID     string  `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	ExpresswayID string  `db:"expressway_id" json:"expressway_id" gorm:"type:varchar(2);"`
	IsReturnTrip string  `db:"is_return_trip" json:"is_return_trip" gorm:"type:varchar(2); comment:'0 = ขาไป 1 =ขากลับ'"`
	SlipID       string  `db:"slip_id" json:"slip_id" gorm:"type:varchar(2);"`
	Price        float64 `db:"price" json:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าทางด่วน' "`
}
