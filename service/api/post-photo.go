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
	"wasa-2024-2024851/service/database"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) addPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	reqToken := getBearerToken(r.Header.Get("Authorization"))
	pathUsername := ps.ByName("userName")
	if pathUsername != reqToken {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// err := r.ParseMultipartForm(10 << 20)
	// if err != nil {
	// 	ctx.Logger.WithError(err).Error("error parsing form")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	file, handler, err := r.FormFile("file")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error when getting file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	println("Uploaded File: ", handler.Filename)
	println("File Size: ", handler.Size)
	println("MIME Header: ", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	id, err := rt.db.AddPhoto(Photo.toDatabase(Photo{Author_ID: pathUsername, Date: time.Now()}))
	if err != nil {
		ctx.Logger.WithError(err).Error("can't add photo to db")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	tempFile, err := os.Create(filepath.Join("/tmp", "/users", pathUsername, "/photos", strconv.FormatInt(id, 10)))
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
	tempFile.Write(fileBytes)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(database.Photo{Author_ID: strconv.FormatInt(id, 10), Date: time.Now()})
}
