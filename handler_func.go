package wgin

import "github.com/gin-gonic/gin"

var (
	errorHandler    ErrorHandler
	responseHandler ResponseHandler
	bindHandler     BindHandler
)

type ErrorHandler func(c *gin.Context, err error)

type ResponseHandler func(c *gin.Context, a any)

type BindHandler func(c *gin.Context, a any) error

type GetBodyHandler[T any] func(*Context[T]) (*T, error)

type BaseHandlerFunc[response any] func(*gin.Context) (*response, error)

type HandlerFunc[request, response any] func(ctx *Context[request]) (*response, error)

// SetErrorHandler 設置預設錯誤處理
func SetErrorHandler(f ErrorHandler) {
	errorHandler = f
}

// SetResponseHandler 設置預設回應處理
func SetResponseHandler(f ResponseHandler) {
	responseHandler = f
}

// SetBindHandler 設置預設 bind
func SetBindHandler(f BindHandler) {
	bindHandler = f
}
