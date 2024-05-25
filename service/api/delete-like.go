package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userUnliker := getBearerToken(r.Header.Get("Authorization"))
	pathPhotoID := ps.ByName("photoId")
	pathUserUnliker := ps.ByName("likeId")
	if pathUserUnliker != userUnliker {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := rt.db.RemoveLike(User.toDatabase(User{ID: userUnliker}), pathPhotoID)
	if err != nil {
		ctx.Logger.WithError(err).Error("could not remove the like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
