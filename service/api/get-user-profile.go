package api

import (
	"encoding/json"
	"net/http"
	"strings"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqUserName := splitToken[1]
	// @TODO: CHECK IF BANNED
	println(reqUserName)
	userName := ps.ByName("userName")
	if userName == "" {
		w.WriteHeader(http.StatusBadRequest)
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
	_ = json.NewEncoder(w).Encode(database.UserProfile{ID: userName, Followings: followings, Followers: followers})
}

func (rt *_router) getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	users, err := rt.db.GetAllUsers()
	if err != nil {
		ctx.Logger.WithError(err).Error("can't retrieve the users")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(users)
}
