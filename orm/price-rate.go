package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type PriceRate struct { // poi ของ รถแต่ละประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PoiID         string  `db:"poi_id" json:"poi_id" gorm:"type:varchar(36);"`
	VehicleTypeID string  `db:"vehicle_type_id" json:"vehicle_type_id" gorm:"type:varchar(36);"`
	Price         float64 `db:"price" json:"price" gorm:"default:null; type:decimal; comment:'ค่าโดยสารของรถแต่ละประเภท(บาท)' "`
}
