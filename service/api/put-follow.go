package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addFollower(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userFollower := getBearerToken(r.Header.Get("Authorization"))
	pathUserToFollow := ps.ByName("userName")
	pathFollowerId := ps.ByName("followerId")
	if userFollower != pathFollowerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	isBanned, err := rt.db.CheckBan(User{ID: pathUserToFollow}.toDatabase(), User{ID: userFollower}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check the ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBanned {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err = rt.db.AddFollower(User{ID: userFollower}.toDatabase(), User{ID: pathUserToFollow}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add the follow")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
