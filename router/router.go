package router

import "net/http"

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
func (r *Router) GET(path string, handler http.Handler) {
	r.handle(http.MethodGet, path, handler)
}
func (r *Router) POST(path string, handler http.Handler) {
	r.handle(http.MethodPost, path, handler)
}
func (r *Router) PUT(path string, handler http.Handler) {
	r.handle(http.MethodPut, path, handler)
}
func (r *Router) DELETE(path string, handler http.Handler) {
	r.handle(http.MethodDelete, path, handler)
}
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if r.PanicHandler != nil {
		defer r.recv(w, req)
	}
}
