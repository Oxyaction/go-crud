package handlers

import (
	mw "github.com/Oxyaction/go-crud/cmd/http/middlewares"
	"github.com/jackc/pgx/v4"
	httprouter "github.com/julienschmidt/httprouter"
)

func AttachHandlers(router *httprouter.Router, db *pgx.Conn) {
	user := &user{}

	middlewares := []mw.Middleware{
		mw.TimeMiddleware,
		mw.AuthMiddleware,
	}

	router.GET("/register", mw.Wrap(middlewares, user.register))
}
