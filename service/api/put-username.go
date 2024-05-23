package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var newUser User
	reqToken := getBearerToken(r.Header.Get("Authorization"))
	pathUsername := ps.ByName("userName")
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("bad body request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if new username is equal to old username, skip db access
	if newUser.ID == pathUsername {
		json.NewEncoder(w).Encode(newUser)
		return
	}
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	exists, err := rt.db.CheckUser(newUser.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't change the username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = rt.db.ChangeUsername(newUser.ID, database.User{ID: pathUsername})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't change the username")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newUser)
}
