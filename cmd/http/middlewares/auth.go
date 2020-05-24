package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(handler httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Authorization: Bearer <token>
		// authHeader, ok := r.Header["Authorization"]
		// if !ok {
		// 	w.WriteHeader(403)
		// 	fmt.Fprintf(w, "Access denied")
		// 	return
		// }
		// fmt.Println("auth header:", authHeader)
		handler(w, r, p)
	})
}
