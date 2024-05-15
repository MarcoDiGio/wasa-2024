package api

// func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// var err error
// var userProfiles []database.UserProfile
// userName := ps.ByName("userName")
// if userName == "" {
// 	w.WriteHeader(http.StatusBadRequest)
// 	return
// }

// userProfiles, err = rt.db.GetUserProfile()
// if err != nil {
// 	if len(userProfiles) == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}
// 	w.WriteHeader(http.StatusInternalServerError)
// 	return
// }

// w.Header().Set("Content-Type", "application/json")
// _ = json.NewEncoder(w).Encode(userProfiles)
// }

// func (rt *_router) getAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
// 	users, err := rt.db.GetAllUsers()
// 	if err != nil {
// 		ctx.Logger.WithError(err).Error("can't retrieve the users")
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	_ = json.NewEncoder(w).Encode(users)
// }
