package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqUsername := getBearerToken(r.Header.Get("Authorization"))
	pathUsername := ps.ByName("userName")
	exists, err := rt.db.CheckUser(User{ID: pathUsername}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when checking if user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	isBanned, err := rt.db.CheckBan(User{ID: pathUsername}.toDatabase(), User{ID: reqUsername}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when checking ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	followers, err := rt.db.GetFollower(User{ID: pathUsername}.toDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	followings, err := rt.db.GetFollowing(User{ID: pathUsername}.toDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photos, err := rt.db.GetAllUserPhotos(User{ID: pathUsername}.toDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(UserProfile{ID: pathUsername, Followings: followings, Followers: followers, Photos: photos}.toDatabase())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
