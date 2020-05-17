package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func TimeMiddleware(handler httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		start := time.Now()
		handler(w, r, p)
		fmt.Printf("Request time: %v\n", time.Since(start))
	})
}
