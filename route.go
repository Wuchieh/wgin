package wgin

import "github.com/gin-gonic/gin"

func UseGet[response any](router gin.IRoutes, path string, handler BaseHandler[response]) {
	router.GET(path, handler.GetHandlers()...)
}

func Get[response any](router gin.IRoutes, path string, handler BaseHandler[response]) {
	router.GET(path, handler.GetHandlers()...)
}

func UseHandle[request, response any](router gin.IRoutes, method, path string, handler Handler[request, response]) {
	router.Handle(method, path, handler.GetHandlers()...)
}

func Handle[request, response any](router gin.IRoutes, method, path string, handler Handler[request, response]) {
	router.Handle(method, path, handler.GetHandlers()...)
}

func UsePOST[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.POST(path, handler.GetHandlers()...)
}

func POST[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.POST(path, handler.GetHandlers()...)
}

func UseDELETE[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.DELETE(path, handler.GetHandlers()...)
}

func DELETE[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.DELETE(path, handler.GetHandlers()...)
}

func UsePATCH[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.PATCH(path, handler.GetHandlers()...)
}

func PATCH[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.PATCH(path, handler.GetHandlers()...)
}

func UsePUT[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.PUT(path, handler.GetHandlers()...)
}

func PUT[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.PUT(path, handler.GetHandlers()...)
}

func UseOPTIONS[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.OPTIONS(path, handler.GetHandlers()...)
}

func OPTIONS[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.OPTIONS(path, handler.GetHandlers()...)
}

func UseHEAD[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.HEAD(path, handler.GetHandlers()...)
}

func HEAD[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.HEAD(path, handler.GetHandlers()...)
}

func UseAny[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.Any(path, handler.GetHandlers()...)
}

func Any[request, response any](router gin.IRoutes, path string, handler Handler[request, response]) {
	router.Any(path, handler.GetHandlers()...)
}
