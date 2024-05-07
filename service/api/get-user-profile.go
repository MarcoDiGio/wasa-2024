package api

import (
	"encoding/json"
	"net/http"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var err error
	var userProfiles []database.UserProfile
	userName := ps.ByName("userName")
	if userName == "" {
		err = err
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userProfiles, err = rt.db.GetUserProfile()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(userProfiles)
}
