package api

import (
	"net/http"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removeComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	pathUsername := ps.ByName("userName")
	pathPhotoId := ps.ByName("photoId")
	pathCommentId := ps.ByName("commentId")
	userDeleter := getBearerToken(r.Header.Get("Authorization"))
	if !isAuthenticated(pathUsername) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if userDeleter != pathUsername {
		isCommenter, err := rt.db.CheckCommentAuthor(pathCommentId, pathPhotoId, User{ID: userDeleter}.toDatabase())
		if err != nil {
			ctx.Logger.WithError(err).Error("could not check if commenter is author")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !isCommenter {
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}
	err := rt.db.RemoveComment(pathCommentId)
	if err != nil {
		ctx.Logger.WithError(err).Error("could not remove comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
