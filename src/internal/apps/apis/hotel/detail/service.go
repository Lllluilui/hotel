package detail

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"hotel/internal/pkg/di"
	"hotel/internal/pkg/models"
)

// Service 服务
type Service struct {

	// ctx 请求上下文
	ctx *gin.Context

	// form 表单
	form *Form

	// hotel 酒店信息
	hotel *models.Hotel
}

// Handler
//
//	@desc: 处理函数
//	@receiver s *Service
//	@return *Response
//	@return int
//	@return error
func (s *Service) Handler() (*Response, int, error) {

	methods := []func() (int, error){
		s.form.Run,
		s.getHotel,
	}

	for _, m := range methods {
		if code, err := m(); err != nil {
			return nil, code, err
		}
	}

	return s.generateResponse(), http.StatusOK, nil
}

// getHotel
//
//	@desc: 获取酒店信息
//	@receiver s *Service
//	@return int
//	@return error
func (s *Service) getHotel() (int, error) {
	var err error

	err = di.Conn(s.ctx).
		Take(&s.hotel).
		Where("id = ?", s.form.ID).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return http.StatusNotFound, err
	}

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

// generateResponse
//
//	@desc: 组合结果
//	@receiver s *Service
//	@return *Response
func (s *Service) generateResponse() *Response {
	hotel := s.hotel

	return &Response{
		ID:        hotel.ID,
		Name:      hotel.Name,
		Address:   hotel.Address,
		Price:     hotel.Price,
		Score:     hotel.Score,
		Brand:     hotel.Brand,
		City:      hotel.City,
		StarName:  hotel.StarName,
		Business:  hotel.Business,
		Pic:       hotel.Pic,
		Longitude: hotel.Longitude,
		Latitude:  hotel.Latitude,
		CreatedAt: hotel.CreatedAt,
		UpdatedAt: hotel.UpdatedAt,
		DeletedAt: hotel.DeletedAt,
	}
}

// NewService
//
//	@desc: 创建服务实例
//	@param ctx *gin.Context
//	@return *Service
func NewService(ctx *gin.Context) *Service {
	return &Service{
		ctx:  ctx,
		form: NewForm(ctx),
	}
}
