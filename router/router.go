package router

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	method string
	path   string
	params []string
}

// Settings holds router-specific configurations like middlewares and SSL settings
type Settings struct {
	middlewares []gin.HandlerFunc
	SSL         bool
	CertKey     string
	CertFile    string
}

type Router struct {
	routes   []*Route
	engine   *gin.Engine
	Settings *Settings
}

// New creates a new Router instance
func New() *Router {
	return &Router{
		engine: gin.Default(),
		Settings: &Settings{
			SSL:      false,
			CertKey:  "",
			CertFile: "",
		},
	}
}
func (r *Router) Get(path string, handler gin.HandlerFunc) {
	r.engine.GET(path, handler)
}
func (r *Router) Post(path string, handler gin.HandlerFunc) {
	r.engine.POST(path, handler)
}
func (r *Router) Put(path string, handler gin.HandlerFunc) {
	r.engine.PUT(path, handler)
}
func (r *Router) Patch(path string, handler gin.HandlerFunc) {
	r.engine.PATCH(path, handler)
}
func (r *Router) Delete(path string, handler gin.HandlerFunc) {
	r.engine.DELETE(path, handler)
}
func (r *Router) Options(path string, handler gin.HandlerFunc) {
	r.engine.OPTIONS(path, handler)
}
func (r *Router) Use(middleware ...gin.HandlerFunc) {
	r.engine.Use(middleware...)
}
func (r *Router) Run(addr string) error {
	if r.Settings.SSL {
		return r.engine.RunTLS(addr, r.Settings.CertFile, r.Settings.CertKey)
	}
	return r.engine.Run(addr)
}
