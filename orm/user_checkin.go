package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type UserCheckIn struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserID           string `db:"user_id" json:"user_id" gorm:"type:varchar(36);"`
	CheckInDatetime  string `db:"check_in_date_time" json:"check_in_date_time" gorm:"type:datetime;"`
	CheckOutDatetime string `db:"check_out_date_time" json:"check_out_date_time" gorm:"type:datetime;"`
	CounterID        string `db:"counter_id" json:"counter_id" gorm:"type:varchar(36);"`
}
