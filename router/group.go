package router

import "net/http"

// Groups
type Group struct {
	Prefix      string
	Middlewares []func(http.Handler) http.Handler
	Router      *Router
}

func (r *Router) Group(prefix string) *Group {
	return &Group{
		Prefix: prefix,
		Router: r,
	}
}

func (g *Group) Use(middleware func(http.Handler) http.Handler) {
	g.Middlewares = append(g.Middlewares, middleware)
}

func (g *Group) AddRoute(method, path string, handler http.Handler) {
	fullPath := g.Prefix + path
	for i := len(g.Middlewares) - 1; i >= 0; i-- {
		handler = g.Middlewares[i](handler)
	}
	g.Router.handle(method, fullPath, handler)
}

func (g *Group) Get(path string, handler http.Handler) {
	g.AddRoute(http.MethodGet, path, handler)
}

func (g *Group) Post(path string, handler http.Handler) {
	g.AddRoute(http.MethodPost, path, handler)
}

func (g *Group) Put(path string, handler http.Handler) {
	g.AddRoute(http.MethodPut, path, handler)
}

func (g *Group) Delete(path string, handler http.Handler) {
	g.AddRoute(http.MethodDelete, path, handler)
}
