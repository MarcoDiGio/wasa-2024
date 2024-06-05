package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userBanner := getBearerToken(r.Header.Get("Authorization"))
	pathBannerId := ps.ByName("userName")
	pathUnbannedId := ps.ByName("bannedId")
	if userBanner != pathBannerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := rt.db.RemoveBan(database.User{ID: userBanner}, database.User{ID: pathUnbannedId})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't remove the ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
