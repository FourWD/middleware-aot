package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type FuelBill struct { //no CRUD
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	SourceID string `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`

	RefuelDateTime string  `db:"refuel_datetime"  json:"refuel_datetime" gorm:"default:null; type:varchar(50); comment:'วันที่เติมน้ำมัน' "`
	VehicleID      string  `db:"vehicle_id" json:"vehicle_id" gorm:"default:null; type:varchar(36); "`
	DriverID       string  `db:"driver_id" json:"driver_id" gorm:"default:null; type:varchar(36); "`
	FuelID         string  `db:"fuel_id" json:"fuel_id" gorm:"type:varchar(2);"`
	Litre          float64 `db:"litre" json:"litre" gorm:"default:null; type:decimal; comment:'จำนวนน้ำมันที่เติม' "`
	Price          float64 `db:"price" json:"price" gorm:"default:null; type:decimal; comment:'ราคาค่าน้ำมัน' "`
}