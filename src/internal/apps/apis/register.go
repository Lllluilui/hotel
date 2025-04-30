package apis

import (
	"github.com/gin-gonic/gin"

	HotelCreateAPI "hotel/internal/apps/apis/hotel/create"
	HotelDeleteAPI "hotel/internal/apps/apis/hotel/delete"
	HotelDetailAPI "hotel/internal/apps/apis/hotel/detail"
	HotelListAPI "hotel/internal/apps/apis/hotel/list"
	HotelUpdateAPI "hotel/internal/apps/apis/hotel/update"
)

// Register 路由注册器
type Register struct{}

// RegisterAndRun
//
//	@desc: 注册并启动
//	@receiver r *Register
//	@return error
func (r *Register) RegisterAndRun() error {
	hotelCreateRoute := new(HotelCreateAPI.API).Route()
	hotelDetailRoute := new(HotelDetailAPI.API).Route()
	hotelDeleteRoute := new(HotelDeleteAPI.API).Route()
	hotelListRoute := new(HotelListAPI.API).Route()
	hotelUpdateRoute := new(HotelUpdateAPI.API).Route()

	s := gin.Default()

	s.POST(hotelCreateRoute.Path, hotelCreateRoute.Handlers...)
	s.GET(hotelDetailRoute.Path, hotelDetailRoute.Handlers...)
	s.DELETE(hotelDeleteRoute.Path, hotelDeleteRoute.Handlers...)
	s.GET(hotelListRoute.Path, hotelListRoute.Handlers...)
	s.PUT(hotelUpdateRoute.Path, hotelUpdateRoute.Handlers...)

	return s.Run()
}
