package detail

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hotel/internal/pkg/server/route"
)

// Form 表单
type Form struct {
	// BaseForm 基础表单
	*route.BaseForm

	// Request 请求体
	*Request
}

// Run
//
//	@desc: 运行
//	@receiver f *Form
//	@return error
func (f *Form) Run() (int, error) {
	var err error

	if err = f.BindData(); err != nil {
		return http.StatusInternalServerError, err
	}

	if err = f.Validate(); err != nil {
		return http.StatusBadRequest, err
	}

	return http.StatusOK, nil
}

// BindData
//
//	@desc: 绑定数据
//	@receiver f *Form
//	@return error
func (f *Form) BindData() error {
	if err := f.BindDataWithURI(f.Request); err != nil {
		return err
	}

	return nil
}

// Validate
//
//	@desc: 验证数据
//	@receiver f *Form
//	@return error
func (f *Form) Validate() error {

	return nil
}

// NewForm
//
//	@desc: 创建表单
//	@param ctx *gin.Context
//	@return *Form
func NewForm(ctx *gin.Context) *Form {
	return &Form{
		BaseForm: route.NewBaseForm(ctx),
		Request:  new(Request),
	}
}
