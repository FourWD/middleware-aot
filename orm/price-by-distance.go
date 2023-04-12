package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type PriceByDistance struct { // poi ของ รถแต่ละประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Distance               float64 `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`
	PriceByVehicleTypeID01 float64 `db:"price_by_vehicle_type_id_01" json:"price_by_vehicle_type_id_01" gorm:"default:null; type:decimal(10,4); comment:'ค่าโดยสารของรถแต่ละประเภท1(บาท)' "`
	PriceByVehicleTypeID02 float64 `db:"price_by_vehicle_type_id_02" json:"price_by_vehicle_type_id_02" gorm:"default:null; type:decimal(10,4); comment:'ค่าโดยสารของรถแต่ละประเภท2(บาท)' "`
	PriceByVehicleTypeID03 float64 `db:"price_by_vehicle_type_id_03" json:"price_by_vehicle_type_id_03" gorm:"default:null; type:decimal(10,4); comment:'ค่าโดยสารของรถแต่ละประเภท3(บาท)' "`
	PriceByVehicleTypeID04 float64 `db:"price_by_vehicle_type_id_04" json:"price_by_vehicle_type_id_04" gorm:"default:null; type:decimal(10,4); comment:'ค่าโดยสารของรถแต่ละประเภท4(บาท)' "`
	PriceByVehicleTypeID05 float64 `db:"price_by_vehicle_type_id_05" json:"price_by_vehicle_type_id_05" gorm:"default:null; type:decimal(10,4); comment:'ค่าโดยสารของรถแต่ละประเภท5(บาท)' "`
	PriceByVehicleTypeID06 float64 `db:"price_by_vehicle_type_id_06" json:"price_by_vehicle_type_id_06" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID07 float64 `db:"price_by_vehicle_type_id_07" json:"price_by_vehicle_type_id_07" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID08 float64 `db:"price_by_vehicle_type_id_08" json:"price_by_vehicle_type_id_08" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID09 float64 `db:"price_by_vehicle_type_id_09" json:"price_by_vehicle_type_id_09" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID10 float64 `db:"price_by_vehicle_type_id_10" json:"price_by_vehicle_type_id_10" gorm:"default:null; type:decimal(10,4); comment:'' "`
}
