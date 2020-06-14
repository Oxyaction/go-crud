package handlers

import (
	mw "github.com/Oxyaction/go-crud/cmd/http/middlewares"
	"github.com/Oxyaction/go-crud/internal/platform/db"
	httprouter "github.com/julienschmidt/httprouter"
)

func AttachHandlers(router *httprouter.Router, db *db.DB) {
	user := NewHandler(db)

	middlewares := []mw.Middleware{
		mw.TimeMiddleware,
		mw.AuthMiddleware,
	}

	router.POST("/register", mw.Wrap(middlewares, user.register))
}
