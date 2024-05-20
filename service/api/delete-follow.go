package api

import (
	"net/http"
	"strings"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteFollower(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	userFollower := splitToken[1]
	pathUserToUnfollow := ps.ByName("userName")
	pathFollowerId := ps.ByName("followerId")
	if userFollower != pathFollowerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// @TODO: CHECK IF BANNED
	err := rt.db.RemoveFollower(database.User{ID: userFollower}, database.User{ID: pathUserToUnfollow})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't remove the follow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
