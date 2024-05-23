package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userName := ps.ByName("userName")
	var user = User{ID: userName}
	photos, err := rt.db.GetAllUserPhoto(user.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't retrieve the photos")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(photos)
}
