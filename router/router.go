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
		ServerMux:   http.NewServeMux(),
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
