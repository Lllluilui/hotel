package models

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTbHotel = "hotel"

// Hotel 酒店信息
type Hotel struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;comment:酒店id" json:"id"`
	Name      string         `gorm:"column:name;type:varchar(255);not null;comment:酒店名称" json:"name"`
	Address   string         `gorm:"column:address;type:varchar(255);not null;comment:酒店地址" json:"address"`
	Price     int32          `gorm:"column:price;type:int;not null;comment:酒店价格" json:"price"`
	Score     int32          `gorm:"column:score;type:int;not null;comment:酒店评分" json:"score"`
	Brand     string         `gorm:"column:brand;type:varchar(32);not null;comment:酒店品牌" json:"brand"`
	City      string         `gorm:"column:city;type:varchar(32);not null;comment:所在城市" json:"city"`
	StarName  string         `gorm:"column:star_name;type:varchar(16);comment:酒店星级，1星到5星，1钻到5钻" json:"star_name"`
	Business  string         `gorm:"column:business;type:varchar(255);comment:商圈" json:"business"`
	Latitude  string         `gorm:"column:latitude;type:varchar(32);not null;comment:纬度" json:"latitude"`
	Longitude string         `gorm:"column:longitude;type:varchar(32);not null;comment:经度" json:"longitude"`
	Pic       string         `gorm:"column:pic;type:varchar(255);comment:酒店图片" json:"pic"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;comment:创建时间" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;comment:更新时间" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;comment:删除时间" json:"deleted_at"`
}

// TableName TbHotel's table name
func (*Hotel) TableName() string {
	return TableNameTbHotel
}
