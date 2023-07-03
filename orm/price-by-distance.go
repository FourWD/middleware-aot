package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type PriceByDistance struct { // poi ของ รถแต่ละประเภท
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	FleetClientID string `db:"fleet_client_id" json:"fleet_client_id" gorm:"type:varchar(10);"`

	Distance float64 `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`

	PriceByVehicleModelID01 float64 `db:"price_by_vehicle_model_id_01" json:"price_by_vehicle_model_id_01" gorm:"default:null; type:decimal(10,4); comment:'BMW' "`
	PriceByVehicleModelID02 float64 `db:"price_by_vehicle_model_id_02" json:"price_by_vehicle_model_id_02" gorm:"default:null; type:decimal(10,4); comment:'BENZ' "`
	PriceByVehicleModelID03 float64 `db:"price_by_vehicle_model_id_03" json:"price_by_vehicle_model_id_03" gorm:"default:null; type:decimal(10,4); comment:'Toyota Camry' "`
	PriceByVehicleModelID04 float64 `db:"price_by_vehicle_model_id_04" json:"price_by_vehicle_model_id_04" gorm:"default:null; type:decimal(10,4); comment:'Toyota Commuter' "`
	PriceByVehicleModelID05 float64 `db:"price_by_vehicle_model_id_05" json:"price_by_vehicle_model_id_05" gorm:"default:null; type:decimal(10,4); comment:'Isuzu Mu-x' "`
	PriceByVehicleModelID06 float64 `db:"price_by_vehicle_model_id_06" json:"price_by_vehicle_model_id_06" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleModelID07 float64 `db:"price_by_vehicle_model_id_07" json:"price_by_vehicle_model_id_07" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleModelID08 float64 `db:"price_by_vehicle_model_id_08" json:"price_by_vehicle_model_id_08" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleModelID09 float64 `db:"price_by_vehicle_model_id_09" json:"price_by_vehicle_model_id_09" gorm:"default:null; type:decimal(10,4); comment:'' "`
	PriceByVehicleModelID10 float64 `db:"price_by_vehicle_model_id_10" json:"price_by_vehicle_model_id_10" gorm:"default:null; type:decimal(10,4); comment:'' "`
}
