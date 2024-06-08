package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userBanner := getBearerToken(r.Header.Get("Authorization"))
	pathBannerId := ps.ByName("userName")
	pathBannedId := ps.ByName("bannedId")
	if userBanner != pathBannerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Banning implies removing follower
	err := rt.db.RemoveFollower(User{ID: pathBannerId}.toDatabase(), User{ID: pathBannedId}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't remove the follow (in put-ban.go)")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = rt.db.AddBan(User{ID: pathBannerId}.toDatabase(), User{ID: pathBannedId}.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add the ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
