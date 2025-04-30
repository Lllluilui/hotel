package create

import (
	"fmt"
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
	hotel []*models.Hotel

	// total 总数
	total int64
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
		s.getHotelList,
	}

	for _, m := range methods {
		if code, err := m(); err != nil {
			return nil, code, err
		}
	}

	return s.generateResponse(), http.StatusOK, nil
}

// buildQuery
//
//	@desc: 构建查询条件
//	@receiver s *Service
//	@return *gorm.DB
func (s *Service) buildQuery() *gorm.DB {
	mode := di.Conn(s.ctx).Model(new(models.Hotel))

	s.buildNameQuery(mode)
	s.buildScoreQuery(mode)
	s.buildBrandQuery(mode)
	s.buildCityQuery(mode)
	s.buildStarNameQuery(mode)
	s.buildBusinessQuery(mode)

	return mode
}

// buildNameQuery
//
//	@desc: 构建酒店名称查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildNameQuery(model *gorm.DB) *gorm.DB {
	if len(s.form.Name) > 0 {
		model.Where("name like ?", fmt.Sprintf("%%%s%", s.form.Name))
	}

	return model
}

// buildScoreQuery
//
//	@desc: 构建酒店评分查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildScoreQuery(model *gorm.DB) *gorm.DB {
	if s.form.Score > 0 {
		model.Where("score > ?", s.form.Score)
	}

	return model
}

// buildBrandQuery
//
//	@desc: 构建酒店品牌查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildBrandQuery(model *gorm.DB) *gorm.DB {
	if len(s.form.Brand) > 0 {
		model.Where("brand = ?", s.form.Brand)
	}

	return model
}

// buildCityQuery
//
//	@desc: 构建酒店城市查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildCityQuery(model *gorm.DB) *gorm.DB {
	if len(s.form.City) > 0 {
		model.Where("city = ?", s.form.City)
	}

	return model
}

// buildStarNameQuery
//
//	@desc: 构建酒店星级查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildStarNameQuery(model *gorm.DB) *gorm.DB {
	if len(s.form.StarName) > 0 {
		model.Where("star_name = ?", s.form.StarName)
	}

	return model
}

// buildBusinessQuery
//
//	@desc: 构建酒店商圈查询条件
//	@receiver s *Service
//	@param model *gorm.DB
//	@return *gorm.DB
func (s *Service) buildBusinessQuery(model *gorm.DB) *gorm.DB {
	if len(s.form.Business) > 0 {
		model.Where("business = ?", s.form.Business)
	}

	return model
}

// getHotelList
//
//	@desc: 获取酒店列表
//	@receiver s *Service
//	@return int
//	@return error
func (s *Service) getHotelList() (int, error) {
	var err error
	queryModel := s.buildQuery()

	err = queryModel.Count(&s.total).Error
	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = queryModel.
		Offset((s.form.PageMark - 1) * s.form.PageSize).
		Limit(s.form.PageSize).
		Find(&s.hotel).
		Error

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
	resp := new(Response)

	resp.Total = s.total

	for _, hotel := range s.hotel {
		h := &Hotel{
			// ID ID（唯一标识）
			ID:        hotel.ID,
			Name:      hotel.Name,
			Address:   hotel.Address,
			Price:     hotel.Price,
			Score:     hotel.Score,
			Brand:     hotel.Brand,
			City:      hotel.City,
			StarName:  hotel.StarName,
			Business:  hotel.Business,
			Latitude:  hotel.Latitude,
			Longitude: hotel.Longitude,
			Pic:       hotel.Pic,
			CreatedAt: hotel.CreatedAt,
			UpdatedAt: hotel.UpdatedAt,
			DeletedAt: hotel.DeletedAt,
		}

		resp.List = append(resp.List, h)
	}

	return resp
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
