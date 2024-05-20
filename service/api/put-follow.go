package api

import (
	"net/http"
	"strings"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addFollower(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	userFollower := splitToken[1]
	pathUserToFollow := ps.ByName("userName")
	pathFollowerId := ps.ByName("followerId")
	if userFollower != pathFollowerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// @TODO: CHECK IF BANNED
	err := rt.db.AddFollower(database.User{ID: userFollower}, database.User{ID: pathUserToFollow})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add the follow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
