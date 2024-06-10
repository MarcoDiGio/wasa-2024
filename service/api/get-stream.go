package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	pathUsername := ps.ByName("userName")
	reqUserName := getBearerToken(r.Header.Get("Authorization"))
	if reqUserName != pathUsername || !isAuthenticated(reqUserName) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var photos = make([]FinalPhoto, 0)
	followPhotos, err := rt.db.GetAllFollowingPhotos(User{ID: reqUserName}.toDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for j, photo := range followPhotos {
		if j >= 5 {
			break
		}
		photos = append(photos, FinalPhoto(photo))
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
