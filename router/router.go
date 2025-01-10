package router

import (
	"github.com/gin-gonic/gin"
)

type Certs struct {
	SSLKey  string
	SSLFile string
}
type Config struct {
	PORT  string
	SSL   bool
	Certs Certs
}
type Route struct {
	Method string
	Path   string
	Params []string
}
type Router struct {
	engine *gin.Engine
	config *Config
}

// New creates a new Router instance
//
//	router := router.New(&router.Config{
//		PORT: ":3000",
//		SSL: true,
//		Certs: Certs{
//			SSLFile: "cert.pem",
//			SSLKey: "key.pem",
//		},
//	})
func New(config *Config) *Router {
	return &Router{
		engine: gin.New(),
		config: config,
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

// starts the router the user inputed Port
//
//	err := router.Run(":3000")
func (r *Router) Run(addr string) error {
	gin.SetMode(gin.ReleaseMode)
	if r.config.SSL {
		return r.engine.RunTLS(addr, r.config.Certs.SSLFile, r.config.Certs.SSLKey)
	}
	return r.engine.Run(addr)
}

//Oldddd Code
// type Route struct {
// 	method string
// 	path   string
// 	params []string
// }

// type Settings struct {
// 	middlewares []gin.HandlerFunc
// 	SSL         bool
// 	CertKey     string
// 	CertFile    string
// }

// type Router struct {
// 	routes   []*Route
// 	engine   *gin.Engine
// 	Settings *Settings
// }

// // New creates a new Router instance
// func New() *Router {
// 	return &Router{
// 		engine: gin.Default(),
// 		Settings: &Settings{
// 			SSL:      false,
// 			CertKey:  "",
// 			CertFile: "",
// 		},
// 	}
// }
// func (r *Router) Get(path string, handler gin.HandlerFunc) {
// 	r.engine.GET(path, handler)
// }
// func (r *Router) Post(path string, handler gin.HandlerFunc) {
// 	r.engine.POST(path, handler)
// }
// func (r *Router) Put(path string, handler gin.HandlerFunc) {
// 	r.engine.PUT(path, handler)
// }
// func (r *Router) Patch(path string, handler gin.HandlerFunc) {
// 	r.engine.PATCH(path, handler)
// }
// func (r *Router) Delete(path string, handler gin.HandlerFunc) {
// 	r.engine.DELETE(path, handler)
// }
// func (r *Router) Options(path string, handler gin.HandlerFunc) {
// 	r.engine.OPTIONS(path, handler)
// }
// func (r *Router) Use(middleware ...gin.HandlerFunc) {
// 	r.engine.Use(middleware...)
// }
// func (r *Router) Run(addr string) error {
// 	if r.Settings.SSL {
// 		return r.engine.RunTLS(addr, r.Settings.CertFile, r.Settings.CertKey)
// 	}
// 	return r.engine.Run(addr)
// }
