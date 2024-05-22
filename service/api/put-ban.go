package api

import (
	"net/http"
	"strings"
	"wasa-2024-2024851/service/api/reqcontext"
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	userBanner := splitToken[1]
	pathBannerId := ps.ByName("userName")
	pathBannedId := ps.ByName("bannedId")
	if userBanner != pathBannerId {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := rt.db.AddBan(database.User{ID: pathBannerId}, database.User{ID: pathBannedId})
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add the ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
