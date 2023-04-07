package orm

import (
	"time"

	orm "github.com/HinekoTech/middleware/orm"
)

type User struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	Created_AT time.Time `db:"created_at" json:"created_at" gorm:"type:datetime"`
	Updated_AT time.Time `db:"updated_at" json:"updated_at" gorm:"type:datetime"`
	Deleted_AT time.Time `db:"deleted_at" json:"deleted_at" gorm:"type:datetime"`
	Created_BY string    `db:"created_by" json:"created_by" gorm:"type:varchar(36)"`
	Updated_BY string    `db:"updated_by" json:"updated_by" gorm:"type:varchar(36)"`
	Deleted_BY string    `db:"deleted_by" json:"deleted_by" gorm:"type:varchar(36)"`
	Domain_ID  string    `db:"domain_id" json:"domain_id" gorm:"type:varchar(36)"`
	Username   string    `db:"username" json:"username" gorm:"type:varchar(50)"`
	Password   string    `db:"password" json:"password" gorm:"type:varchar(100)"`
	Firstname  string    `db:"firstname" json:"firstname" gorm:"type:varchar(50)"`
	Lastname   string    `db:"lastname" json:"lastname" gorm:"type:varchar(50)"`
	Nickname   string    `db:"nickname" json:"nickname" gorm:"type:varchar(50)"`
	Birthday   time.Time `db:"birthday" json:"birthday" gorm:"type:datetime"`
	Avatar     string    `db:"avatar" json:"avatar" gorm:"type:varchar(100)"`
	Email      string    `db:"email" json:"email" gorm:"type:varchar(50)"`
	Position   string    `db:"position" json:"position" gorm:"type:varchar(100)"`
}
