package orm

import (
	"time"

	orm "github.com/HinekoTech/middleware/orm"
)

type User struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	Username  string    `db:"username" json:"username" gorm:"type:varchar(50)"`
	Password  string    `db:"password" json:"password" gorm:"type:varchar(100)"`
	Firstname string    `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname  string    `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Nickname  string    `db:"nickname" json:"nickname" gorm:"type:varchar(50)"`
	Birthday  time.Time `db:"birthday" json:"birthday" gorm:"type:datetime"`
	Avatar    string    `db:"avatar" json:"avatar" gorm:"type:varchar(100)"`
	Email     string    `db:"email" json:"email" gorm:"type:varchar(50)"`
	Position  string    `db:"position" json:"position" gorm:"type:varchar(100)"`
}
