package orm

import (
	orm "github.com/FourWD/middleware/orm"
)

type Slip struct {
	ID string `db:"id" json:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SlipNo            string `db:"slip_no"  json:"slip_no" gorm:"type:varchar(50); index"`
	SlipTypeID        string `db:"slip_type_id" json:"slip_type_id" gorm:"type:varchar(2); comment:'call center website etc.'"`
	SlipTypeServiceID string `db:"slip_type_service_id" json:"slip_type_service_id" gorm:"type:varchar(2); comment:'oneway pickup etc.'"`
	SlipVehicleTypeID string `db:"slip_veh_type_id" json:"slip_veh_type_id" gorm:"type:varchar(36); comment:'ประเภทรถตามหน้าตั๋วที่ซื้อ'"`
	IsPickup          bool   `db:"is_pickup" json:"is_pickup" gorm:"default:0; type:tinyint(1); comment:'รับลูกค้าหรือยัง' "`
	IsPickupDateTime  string `db:"is_pickup_datetime" json:"is_pickup_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่ไปรับห' "`

	CounterIP           string `db:"counter_ip" json:"counter_ip" gorm:"type:varchar(20)"`
	BookingBy           string `db:"booking_by" json:"booking_by" gorm:"type:varchar(36); comment:'จองโดย'"`
	BookingTypeID       string `db:"booking_type_id" json:"booking_type_id" gorm:"type:varchar(2);"`
	BookingDateTime     string `db:"booking_datetime" json:"booking_datetime" gorm:"default:null; type:varchar(50); comment:'วันที่จอง' "`
	OriginPoiID         string `db:"origin_poi_id" json:"origin_poi_id" gorm:"type:varchar(36);"`
	DesinationPoiID     string `db:"desination_poi_id" json:"desination_poi_id" gorm:"type:varchar(36);"`
	ForceDesinationName string `db:"force_desination_name" json:"force_desination_name" gorm:"type:varchar(150) ; comment:'กรณีที่ไม่มี POI' "`

	Distance        float64 `db:"distance" json:"distance" gorm:"type:decimal(16,4)"`
	PriceRateID     string  `db:"price_rate_id" json:"price_rate_id" gorm:"type:varchar(36);"`
	PromotionID     string  `db:"promotion_id" json:"promotion_id" gorm:"type:varchar(36);"`
	PromotionRef    string  `db:"promotion_ref" json:"promotion_ref" gorm:"type:varchar(36); comment:'กรณี duoslip เก็บเลขที่ slip ถ้าkansaiเก็บ เลขของ kansai' "`
	Price           float64 `db:"price" json:"price" gorm:"default:null; type:decimal(16,4); comment:'ราคาค่าบริการ' "`
	Discount        float64 `db:"discount" json:"discount" gorm:"default:null; type:decimal(16,4); comment:'ส่วนลด' "`
	Wht             float64 `db:"wht" json:"wht" gorm:"default:null; type:decimal(16,4); comment:'ภาษีหัก ที่จ่าย' "`
	Vat             float64 `db:"vat" json:"vat" gorm:"default:null; type:decimal(16,4); comment:'ภาษี' "`
	NetPrice        float64 `db:"netprice" json:"netprice" gorm:"default:null; type:decimal(16,4); comment:'รวมราคาค่าบริการทั้งหมด' "`
	IsPaid          bool    `db:"is_paid" json:"is_paid" gorm:"default:0; type:tinyint(1); comment:'ลูกค้าจ่ายหรือยัง' "`
	PaymentTypeID   string  `db:"payment_type_id" json:"payment_type_id" gorm:"type:varchar(2);"`
	PaymentDateTime string  `db:"payment_datetime" json:"payment_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่จ่าย' "`
	IsVoid          bool    `db:"is_void" json:"is_void" gorm:"default:0; type:tinyint(1); comment:'ยกเลิก?' "`
	VoidBy          string  `db:"void_by" json:"void_by" gorm:"type:varchar(36);"`
	VoidDateTime    string  `db:"void_datetime" json:"void_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่ยกเลิก' "`

	IsNewCustomer bool   `db:"is_newcustomer" json:"is_newcustomer" gorm:"type:tinyint(2)"`
	CustomerID    string `db:"customer_id" json:"customer_id" gorm:"type:varchar(36);"`

	CompanyName string `db:"company_name" json:"company_name" gorm:"type:varchar(150)"`
	TaxNo       string `db:"tax_no" json:"tax_no" gorm:"type:varchar(20)"`
	IsHQ        bool   `db:"is_hq" json:"is_hq" gorm:"type:tinyint(2)"`
	Address     string `db:"address" json:"address" gorm:"type:text"`
	Postcode    string `db:"postcode" json:"postcode" gorm:"type:varchar(5)"`
	PhoneNo     string `db:"phone_no" json:"phone_no" gorm:"type:varchar(10)"`
	Fight       string `db:"fight" json:"fight" gorm:"type:varchar(5)"`

	AssignVehicleID string `db:"assign_vehicle_id" json:"assign_vehicle_id" gorm:"type:varchar(36); comment:'รถที่วิ่งงานจริง'"`
	AssignVehicleBy string `db:"assign_vehicle_by" json:"assign_vehicle_by" gorm:"type:varchar(36);"`
	AssignDriverID  string `db:"assign_driver_id" json:"assign_driver_id" gorm:"default:null; type:varchar(36); "`
	ArrivedDateTime string `db:"arrived_datetime" json:"arrived_datetime" gorm:"default:null; type:varchar(50); comment:'วันเวลาที่มาถึง' "`

	IsCompleted bool `db:"is_completed" json:"is_completed" gorm:"default:0; type:tinyint(1); comment:'เสร็จสมบูรณ์' "`
}
