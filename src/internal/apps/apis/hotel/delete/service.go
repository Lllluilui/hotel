package delete

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		s.deleteHotel,
	}

	for _, m := range methods {
		if code, err := m(); err != nil {
			return nil, code, err
		}
	}

	return s.generateResponse(), http.StatusOK, nil
}

// deleteHotel
//
//	@desc: 删除酒店
//	@receiver s *Service
//	@return int
//	@return error
func (s *Service) deleteHotel() (int, error) {
	var err error
	s.hotel = new(models.Hotel)
	s.hotel.ID = s.form.ID

	err = di.Conn(s.ctx).
		Delete(&s.hotel).
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
	return &Response{}
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
