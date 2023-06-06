package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Vehicle struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SourceID          string  `db:"source_id" json:"source_id" gorm:"type:varchar(2);"`
	Code              string  `db:"code"  json:"code" gorm:"type:varchar(50) ; dafault:null ; index"`
	LicensePlate      string  `db:"license_plate" json:"license_plate" gorm:"type:varchar(50) ; dafault:null "`
	VehicleTypeID     string  `db:"vehicle_type_id" json:"vehicle_type_id" gorm:"type:varchar(2);"`
	VehicleBrandID    string  `db:"vehicle_brand_id" json:"vehicle_brand_id" gorm:"type:varchar(2) "`
	VehicleModelID    string  `db:"vehicle_model_id" json:"vehicle_model_id" gorm:"type:varchar(2);"`
	VehicleSubmodelID string  `db:"vehicle_submodel_id" json:"vehicle_submodel_id" gorm:"type:varchar(2); "`
	Latitude          float64 `db:"lat" json:"lat" gorm:"type:decimal(10,6)"`
	Longitude         float64 `db:"long" json:"long" gorm:"type:decimal(10,6)"`
	VehicleStatusID   string  `db:"vehicle_status_id" json:"vehicle_status_id" gorm:"type:varchar(2);"`
	LocationTypeID    string  `db:"location_type_id" json:"location_type_id" gorm:"type:varchar(2);"`
	CurrentDriverID   string  `db:"current_driver_id" json:"current_driver_id" gorm:"default:null; type:varchar(36); "`
}
