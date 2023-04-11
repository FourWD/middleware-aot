package orm

import (
	orm "github.com/HinekoTech/middleware/orm"
)

type Slip struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	BookingNo            string  `db:"booking_no"  json:"booking_no" gorm:"type:varchar(50); index"`
	BookingTypeID        string  `db:"booking_type_id" json:"booking_type_id" gorm:"type:varchar(2);"`
	BookingVehicleTypeID string  `db:"booking_veh_type_id" json:"booking_veh_type_id" gorm:"type:varchar(36);"`
	SlipTypeID           string  `db:"slip_type_id" json:"slip_type_id" gorm:"type:varchar(2);"`
	IsPickup             bool    `db:"is_pickup" json:"is_pickup" gorm:"default:0; type:tinyint(1); comment:'รับลูกค้าหรือยัง' "`
	CounterIP            string  `db:"counter_ip" json:"counter_ip" gorm:"type:varchar(20)"`
	BookingBy            string  `db:"booking_by" json:"booking_by" gorm:"type:varchar(36); comment:'จองโดย'"`
	BookingDateTime      string  `db:"booking_datetime" json:"booking_datetime" gorm:"default:null; type:varchar(50); comment:'วันที่จอง' "`
	PoiID                string  `db:"poi_id" json:"poi_id" gorm:"type:varchar(36);"`
	ForceDesinationName  string  `db:"force_desination_name" json:"force_desination_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`
	Distance             float64 `db:"distance" json:"distance" gorm:"type:decimal"`
	PriceRateID          string  `db:"price_rate_id" json:"price_rate_id" gorm:"type:varchar(36);"`
	Price                float64 `db:"price" json:"price" gorm:"default:null; type:decimal; comment:'ราคาค่าบริการ' "`
	Discount             float64 `db:"discount" json:"discount" gorm:"default:null; type:decimal; comment:'ส่วนลด' "`
	Vat                  float64 `db:"vat" json:"vat" gorm:"default:null; type:decimal; comment:'ภาษี' "`
	NetPrice             float64 `db:"netprice" json:"netprice" gorm:"default:null; type:decimal; comment:'รวมราคาค่าบริการทั้งหมด' "`
	IsPaid               bool    `db:"is_paid" json:"is_paid" gorm:"default:0; type:tinyint(1); comment:'ลูกค้าจ่ายหรือยัง' "`
	PaymentTypeID        string  `db:"payment_type_id" json:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentDateTime      string  `db:"payment_datetime" json:"payment_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่จ่าย' "`
	IsVoid               bool    `db:"is_void" json:"is_void" gorm:"default:0; type:tinyint(1); comment:'ยกเลิก?' "`
	VoidBy               string  `db:"void_by" json:"void_by" gorm:"type:varchar(36);"`
	VoidDateTime         string  `db:"void_datetime" json:"void_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่ยกเลิก' "`

	CustomerID string `db:"customer_id" json:"customer_id" gorm:"type:varchar(36);"`

	AssignVehicleID string `db:"assign_vehicle_id" json:"assign_vehicle_id" gorm:"type:varchar(36);"`
	AssignVehicleBy string `db:"assign_vehicle_by" json:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID  string `db:"assign_driver_id" json:"assign_driver_id" gorm:"default:null; type:varchar(36); "`
	ArrivedDateTime string `db:"arrived_datetime" json:"arrived_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่มาถึง' "`

	IsCompleted bool `db:"is_completed" json:"is_completed" gorm:"default:0; type:tinyint(1); comment:'เสร็จสมบูรณ์' "`
}
