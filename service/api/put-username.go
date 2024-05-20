package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) changeUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var newUser User
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]
	pathUsername := ps.ByName("userName")
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("bad body request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = rt.db.ChangeUsername(newUser.ID, database.User{ID: pathUsername})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't change the username")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}
