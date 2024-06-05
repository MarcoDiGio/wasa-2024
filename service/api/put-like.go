package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	userLiker := getBearerToken(r.Header.Get("Authorization"))
	// pathAuthorID := ps.ByName("userName")
	pathPhotoID := ps.ByName("photoId")
	pathUserLiker := ps.ByName("likeId")
	if pathUserLiker != userLiker {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// @TODO: check if banned by author
	err := rt.db.AddLike(User.toDatabase(User{ID: userLiker}), pathPhotoID)
	if err != nil {
		ctx.Logger.WithError(err).Error("could not add the like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
