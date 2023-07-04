package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type FleetClient struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(10);primary_key;"`
	orm.GormModel

	FleetID string `db:"fleet_id" json:"fleet_id" gorm:"type:varchar(4)"`

	Name     string `db:"name" json:"name" gorm:"type:varchar(50)"`
	RefCode  string `db:"ref_code" json:"ref_code" gorm:"type:varchar(50)"`
	Remark   string `db:"remark" json:"remark" gorm:"type:varchar(200)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
