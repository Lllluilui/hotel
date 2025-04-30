package create

import (
	"time"

	"gorm.io/gorm"
)

// Request 请求
type Request struct {
	// Name 酒店名称
	Name string `json:"name"`

	// Address 地址
	Address string `json:"address"`

	// Price 价格
	Price int32 `json:"price"`

	// Score 评分
	Score int32 `json:"score"`

	// Brand 品牌
	Brand string `json:"brand"`

	// City 城市
	City string `json:"city"`

	// StarName 星级
	StarName string `json:"starName"`

	// Business 商圈
	Business string `json:"business"`

	// Latitude 经度
	Latitude string `json:"latitude"`

	// Longitude 纬度
	Longitude string `json:"longitude"`

	// Pic 图片
	Pic string `json:"pic"`
}

// Response 响应
type Response struct {
	// ID ID（唯一标识）
	ID int64

	// Name 酒店名称
	Name string `json:"name"`

	// Address 地址
	Address string `json:"address"`

	// Price 价格
	Price int32 `json:"price"`

	// Score 评分
	Score int32 `json:"score"`

	// Brand 品牌
	Brand string `json:"brand"`

	// City 城市
	City string `json:"city"`

	// StarName 星级
	StarName string `json:"starName"`

	// Business 商圈
	Business string `json:"business"`

	// Latitude 经度
	Latitude string `json:"latitude"`

	// Longitude 纬度
	Longitude string `json:"longitude"`

	// Pic 图片
	Pic string `json:"pic"`

	// 创建时间
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at"`

	// DeletedAt 删除时间
	DeletedAt gorm.DeletedAt
}
