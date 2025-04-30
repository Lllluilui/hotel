package route

import (
	"github.com/gin-gonic/gin"
)

// Entry 路由表实体
type Entry struct {
	// Path 请求路径
	Path string

	// Method 请求方法
	Method string

	// Handlers 处理请求方法
	Handlers []gin.HandlerFunc
}
