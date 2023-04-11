package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type FavoritePoi struct { // ปลายทางใช้แสดงที่ระบบขายหน้าแรก
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	PoiID    string `db:"poi_id" json:"poi_id" gorm:"type:varchar(36);"`
	RowOrder int    `db:"row_order" json:"row_order" gorm:"type:int"`
}
