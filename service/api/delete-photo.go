package api

import (
	"net/http"
	"os"
	"path/filepath"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) removePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := getBearerToken(r.Header.Get("Authorization"))
	pathUsername := ps.ByName("userName")
	pathPhotoId := ps.ByName("photoId")
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := rt.db.RemovePhoto(pathPhotoId)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't remove the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	path := filepath.Join("/tmp", "/users", pathUsername, "/photos", pathPhotoId)
	err = os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		ctx.Logger.WithError(err).Error("can't remove the photo from filesystem")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
