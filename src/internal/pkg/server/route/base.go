package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// BaseForm 基础表单
type BaseForm struct {
	Ctx *gin.Context
}

// BindDataWithURI
//
//	@desc: 绑定 URI
//	@receiver f *BaseForm
//	@param data any
//	@return error
func (f *BaseForm) BindDataWithURI(data any) error {
	if err := f.Ctx.ShouldBindUri(data); err != nil {
		return err
	}

	return nil
}

// BindDataWithQuery
//
//	@desc: 绑定 query 参数
//	@receiver f *BaseForm
//	@param data any
//	@return error
func (f *BaseForm) BindDataWithQuery(data any) error {
	if err := f.Ctx.ShouldBindQuery(data); err != nil {
		return err
	}

	return nil
}

// BindDataWithHeader
//
//	@desc: 绑定 header
//	@receiver f *BaseForm
//	@param data any
//	@return error
func (f *BaseForm) BindDataWithHeader(data any) error {
	if err := f.Ctx.ShouldBindHeader(data); err != nil {
		return err
	}

	return nil
}

// BindDataWithBody
//
//	@desc: 绑定 body
//	@receiver f *BaseForm
//	@param data any
//	@return error
func (f *BaseForm) BindDataWithBody(data any) error {
	if err := f.Ctx.ShouldBindBodyWith(data, binding.JSON); err != nil {
		return err
	}

	return nil
}

// NewBaseForm
//
//	@desc: 创建一个基础表单
//	@param ctx *gin.Context
//	@return *BaseForm
func NewBaseForm(ctx *gin.Context) *BaseForm {
	return &BaseForm{
		Ctx: ctx,
	}
}
