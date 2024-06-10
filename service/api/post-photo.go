package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"wasa-2024-2024851/service/api/reqcontext"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := getBearerToken(r.Header.Get("Authorization"))
	pathUsername := ps.ByName("userName")
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when getting file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	timeNow := time.Now()
	photoId, err := rt.db.AddPhoto(Photo.toDatabase(Photo{Author_ID: pathUsername, Date: timeNow}))
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add photo to db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tempFile, err := os.Create(filepath.Join("/tmp", "/users", pathUsername, "/photos", strconv.FormatInt(photoId, 10)))
	if err != nil {
		ctx.Logger.WithError(err).Error("can't create image")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write this byte array to our temporary file
	_, err = tempFile.Write(fileBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(FinalPhoto{Photo_ID: strconv.FormatInt(photoId, 10), Author_ID: pathUsername, Date: timeNow, Comments: nil, Likes: nil})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("couldn't convert go values to JSON")
		return
	}
}
