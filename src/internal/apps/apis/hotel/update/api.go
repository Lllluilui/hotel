package create

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hotel/internal/pkg/server/route"
)

// API 接口
type API struct{}

// Invoke
//
//	@desc:
//	@receiver api *API
//	@param ctx *gin.Context
func (api *API) Invoke(ctx *gin.Context) {
	data, code, err := NewService(ctx).Handler()

	if err != nil {
		ctx.JSON(code, route.NewHandlerError(err.Error(), nil, code))
		return
	}

	ctx.JSON(http.StatusOK, route.NewHandlerSuccess(data))
}

// Route
//
//	@desc: 路由
//	@receiver api *API
//	@return route.Entry
func (api *API) Route() route.Entry {
	return route.Entry{
		Path:   "/api/hotel/:hotelId",
		Method: http.MethodPut,
		Handlers: []gin.HandlerFunc{
			api.Invoke,
		},
	}
}
