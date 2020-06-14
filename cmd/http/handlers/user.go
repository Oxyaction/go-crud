package handlers

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/Oxyaction/go-crud/internal/platform/db"
	"github.com/Oxyaction/go-crud/internal/user"
	"github.com/Oxyaction/go-crud/pkg/jsonutil"
	httprouter "github.com/julienschmidt/httprouter"
)

type userHandler struct {
	userManager *user.UserManager
}

func NewHandler(db *db.DB) *userHandler {
	return &userHandler{
		userManager: user.NewManager(db),
	}
}

func (h *userHandler) register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userToCreate user.UserCreate
	code, err := jsonutil.Decode(w, r, &userToCreate)
	if err != nil {
		http.Error(w, err.Error(), code)
		return
	}

	user, err := h.userManager.Register(r.Context(), userToCreate)

	if err != nil {
		log.Error("user register error ", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		log.Error("json marshalization failed", err)
		// log.Printf("user register json encoding error %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(200)
	w.Write(userJson)
}
