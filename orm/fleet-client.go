package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type FleetClient struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(2);primary_key;"`
	orm.GormModel

	FleetID string `db:"fleet_id" json:"fleet_id" gorm:"type:varchar(36)"`

	Name     string `db:"name" json:"name" gorm:"type:varchar(50)"`
	Remark   string `db:"remark" json:"remark" gorm:"type:varchar(100)"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}