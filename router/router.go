package router

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Router struct {
	*http.ServeMux
	middlewares []Middleware
}

// NewRouter creates a new Router instance with optinal middlewares
func NewRouter(middlewares ...Middleware) *Router {
	return &Router{
		ServeMux:    http.NewServeMux(),
		middlewares: middlewares,
	}
}

// Use adds middleware to Router
func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

func (r *Router) AddMiddlewares(middleware ...Middleware) {
	r.middlewares = append(r.middlewares, middleware...)
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	for _, middleware := range r.middlewares {
		handler = middleware(handler)
	}

	r.ServeMux.Handle(pattern, handler)
}

func (r *Router) HandleFunc(pattern string, handlerFunc http.HandlerFunc) {
	var handler http.Handler
	for _, mw := range r.middlewares {
		handler = mw(http.HandlerFunc(handlerFunc))
	}
	r.ServeMux.Handle(pattern, handler)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.ServeMux.ServeHTTP(w, r)
}
