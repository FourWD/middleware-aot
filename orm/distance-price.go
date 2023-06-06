package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type DistancePrice struct { //
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	Month int `db:"imonth" json:"imonth" gorm:"type:int(2);"`
	Year  int `db:"iyear" json:"iyear" gorm:"type:int(4);"`

	PriceByVehicleTypeID01 float64 `db:"price_by_vehicle_type_id_01" json:"price_by_vehicle_type_id_01" gorm:"default:null; type:decimal(10,4); comment:'BMW' "`
	PriceByVehicleTypeID02 float64 `db:"price_by_vehicle_type_id_02" json:"price_by_vehicle_type_id_02" gorm:"default:null; type:decimal(10,4); comment:'BENZ' "`
	PriceByVehicleTypeID03 float64 `db:"price_by_vehicle_type_id_03" json:"price_by_vehicle_type_id_03" gorm:"default:null; type:decimal(10,4); comment:'Toyota Camry' "`
	PriceByVehicleTypeID04 float64 `db:"price_by_vehicle_type_id_04" json:"price_by_vehicle_type_id_04" gorm:"default:null; type:decimal(10,4); comment:'Toyota Commuter' "`
	PriceByVehicleTypeID05 float64 `db:"price_by_vehicle_type_id_05" json:"price_by_vehicle_type_id_05" gorm:"default:null; type:decimal(10,4); comment:'Isuzu Mu-x' "`
	PriceByVehicleTypeID06 float64 `db:"price_by_vehicle_type_id_06" json:"price_by_vehicle_type_id_06" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID07 float64 `db:"price_by_vehicle_type_id_07" json:"price_by_vehicle_type_id_07" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID08 float64 `db:"price_by_vehicle_type_id_08" json:"price_by_vehicle_type_id_08" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID09 float64 `db:"price_by_vehicle_type_id_09" json:"price_by_vehicle_type_id_09" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleTypeID10 float64 `db:"price_by_vehicle_type_id_10" json:"price_by_vehicle_type_id_10" gorm:"default:null; type:decimal(10,4); comment:'' "`
}
