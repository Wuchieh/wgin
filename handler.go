package wgin

import "github.com/gin-gonic/gin"

// BaseHandler 沒有 request
type BaseHandler[response any] struct {
	Middleware      []gin.HandlerFunc
	ErrorHandler    ErrorHandler
	ResponseHandler ResponseHandler
	HandlerFunc     BaseHandlerFunc[response]
}

// GetHandlers 取得 []gin.HandlerFunc
func (h BaseHandler[T]) GetHandlers() []gin.HandlerFunc {
	m := h.Middleware
	return append(m, h.Handler())
}

// Handler 處理 HandlerFunc 後 依據條件處理 ErrorHandler ResponseHandler
func (h BaseHandler[T]) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp *T
		var err error

		if h.HandlerFunc != nil {
			resp, err = h.HandlerFunc(c)
		}

		if err != nil {
			if h.ErrorHandler != nil {
				h.ErrorHandler(c, err)
			} else if errorHandler != nil {
				errorHandler(c, err)
			}
			return
		}

		if h.ResponseHandler != nil {
			h.ResponseHandler(c, resp)
		} else if responseHandler != nil {
			responseHandler(c, resp)
		}
	}
}

// SetErrorHandler 設置錯誤處理函數
func (h BaseHandler[T]) SetErrorHandler(f ErrorHandler) BaseHandler[T] {
	h.ErrorHandler = f
	return h
}

// SetResponseHandler 設置回應處理函數
func (h BaseHandler[T]) SetResponseHandler(f ResponseHandler) BaseHandler[T] {
	h.ResponseHandler = f
	return h
}

// AddMiddleware 添加中間件
func (h BaseHandler[T]) AddMiddleware(f gin.HandlerFunc, fs ...gin.HandlerFunc) BaseHandler[T] {
	m := h.Middleware
	m = append(m, f)
	if len(fs) > 0 {
		m = append(m, fs...)
	}
	h.Middleware = m
	return h
}

func NewBaseHandler[response any](f BaseHandlerFunc[response]) BaseHandler[response] {
	return BaseHandler[response]{
		HandlerFunc: f,
	}
}

// Handler 有 request 和 response
type Handler[request, response any] struct {
	BaseHandler[response]
	GetBodyHandler GetBodyHandler[request]
	HandlerFunc    HandlerFunc[request, response]
}

// SetErrorHandler 設置錯誤處理函數
func (h Handler[T, E]) SetErrorHandler(f ErrorHandler) Handler[T, E] {
	h.BaseHandler = h.BaseHandler.SetErrorHandler(f)
	return h
}

// SetResponseHandler 設置回應處理函數
func (h Handler[T, E]) SetResponseHandler(f ResponseHandler) Handler[T, E] {
	h.BaseHandler = h.BaseHandler.SetResponseHandler(f)
	return h
}

// AddMiddleware 添加中間件
func (h Handler[T, E]) AddMiddleware(f gin.HandlerFunc, fs ...gin.HandlerFunc) Handler[T, E] {
	h.BaseHandler = h.BaseHandler.AddMiddleware(f, fs...)
	return h
}

// Handler 取得 gin.HandlerFunc
func (h Handler[T, E]) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var resp *E
		var err error

		if h.HandlerFunc != nil {
			resp, err = h.HandlerFunc(h.getContext(c))
		}

		if err != nil {
			if h.ErrorHandler != nil {
				h.ErrorHandler(c, err)
			} else if errorHandler != nil {
				errorHandler(c, err)
			}
			return
		}

		if h.ResponseHandler != nil {
			h.ResponseHandler(c, resp)
		} else if responseHandler != nil {
			responseHandler(c, resp)
		}
	}
}

// GetHandlers 取得 []gin.HandlerFunc
func (h Handler[T, E]) GetHandlers() []gin.HandlerFunc {
	m := h.Middleware
	return append(m, h.Handler())
}

func (h Handler[T, E]) getContext(c *gin.Context) *Context[T] {
	return &Context[T]{
		BaseContext:    c,
		getBodyHandler: h.GetBodyHandler,
	}
}

func NewHandler[request, response any](f HandlerFunc[request, response]) Handler[request, response] {
	return Handler[request, response]{
		HandlerFunc: f,
	}
}
