package router

import (
	"net/http"
)

//	func main() {
//	    router := router.New()
//	    fmt.Printf(http.ListenAndServe(":8080", router))
//	}
type Router struct {
	NotFound         http.Handler
	MethodNotAllowed http.Handler
	PanicHandler     func(http.ResponseWriter, *http.Request, interface{})
	routes           map[string]map[string]http.Handler
}

func New() *Router {
	return &Router{
		routes: make(map[string]map[string]http.Handler),
	}
}
func (r *Router) recv(w http.ResponseWriter, req *http.Request) {
	if rcv := recover(); rcv != nil {
		r.PanicHandler(w, req, rcv)
	}
}
func (r *Router) handle(method, path string, handler http.Handler) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.Handler)
	}
	r.routes[method][path] = handler
}
func (r *Router) Get(path string, handler http.Handler) {
	r.handle(http.MethodGet, path, handler)
}
func (r *Router) Post(path string, handler http.Handler) {
	r.handle(http.MethodPost, path, handler)
}
func (r *Router) Put(path string, handler http.Handler) {
	r.handle(http.MethodPut, path, handler)
}
func (r *Router) Delete(path string, handler http.Handler) {
	r.handle(http.MethodDelete, path, handler)
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}
	methodRoutes := r.routes[req.Method]
	if methodRoutes != nil {
		if handler, ok := methodRoutes[req.URL.Path]; ok {
			handler.ServeHTTP(w, req)
			return
		}
	}
	for _, routes := range r.routes {
		if _, ok := routes[req.URL.Path]; ok {
			if r.MethodNotAllowed != nil {
				r.MethodNotAllowed.ServeHTTP(w, req)
			} else {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}
	if r.NotFound != nil {
		r.NotFound.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}
