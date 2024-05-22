package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userName := ps.ByName("userName")
	photos, err := rt.db.GetAllUserPhoto(database.User{ID: userName})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't retrieve the photos")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)
}
