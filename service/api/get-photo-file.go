package api

import (
	"net/http"
	"path/filepath"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	http.ServeFile(w, r, filepath.Join("/tmp", "/users", ps.ByName("userName"), "/photos", ps.ByName("photoId")))
}
