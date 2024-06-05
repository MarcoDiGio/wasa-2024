package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// pathUsername := ps.ByName("userName")
	pathPhotoId := ps.ByName("photoId")
	reqToken := getBearerToken(r.Header.Get("Authorization"))
	var comment Comment
	user := User{ID: reqToken}
	if !isAuthenticated(reqToken) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		ctx.Logger.WithError(err).Error("could not decode the body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// @TODO: CHECK IF COMMENTER IS BANNED
	commentId, err := rt.db.AddComment(pathPhotoId, user.toDatabase(), comment.toDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("could not add the comment")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	comment.Photo_ID = pathPhotoId
	comment.User_ID = user.ID
	comment.Comment_ID = strconv.FormatInt(commentId, 10)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
