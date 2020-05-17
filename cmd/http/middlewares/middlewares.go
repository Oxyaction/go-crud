package middlewares

import (
	"github.com/julienschmidt/httprouter"
)

type Middleware func(handler httprouter.Handle) httprouter.Handle

func Wrap(middlewares []Middleware, handler httprouter.Handle) httprouter.Handle {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
