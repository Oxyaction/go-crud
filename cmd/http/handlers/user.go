package handlers

import (
	"fmt"
	"net/http"

	httprouter "github.com/julienschmidt/httprouter"
)

type user struct {
}

func (u *user) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello world")
}
