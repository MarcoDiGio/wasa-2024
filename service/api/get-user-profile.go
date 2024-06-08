package api

import (
	"encoding/json"
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	reqUserName := getBearerToken(r.Header.Get("Authorization"))
	userName := ps.ByName("userName")
	exists, err := rt.db.CheckUser(database.User{ID: userName})
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when checking if user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	isBanned, err := rt.db.CheckBan(database.User{ID: userName}, database.User{ID: reqUserName})
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when checking ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	followers, err := rt.db.GetFollower(database.User{ID: userName})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	followings, err := rt.db.GetFollowing(database.User{ID: userName})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(database.UserProfile{ID: userName, Followings: followings, Followers: followers})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}

// REMEMBER TO DELETE THIS
func (rt *_router) getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	users, err := rt.db.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("can't retrieve the users")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
