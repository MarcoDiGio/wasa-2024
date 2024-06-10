package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	pathUsername := ps.ByName("userName")
	var user = User{ID: pathUsername}
	if !isAuthenticated(user.ID) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	photos, err := rt.db.GetAllUserPhotos(user.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't retrieve the photos")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
