package api

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
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
		err := json.NewEncoder(w).Encode(newUser)
		if err != nil {
			ctx.Logger.WithError(err).Error("bad body request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		return
	}
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusForbidden)
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
	oldPath := filepath.Join("/tmp", "/users", pathUsername)
	newPath := filepath.Join("/tmp", "/users", newUser.ID)
	err = os.Rename(oldPath, newPath)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't rename directory")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
