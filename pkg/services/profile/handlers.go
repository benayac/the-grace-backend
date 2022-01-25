package profile

import (
	"encoding/json"
	"log"
	"net/http"
	"thegrace/pkg/db"
	"thegrace/pkg/middleware"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET PROFILE][REQUEST]")

	email, err := middleware.ParseAuth(r.Header, middleware.KeyClient)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[GET PROFILE][ERROR] Parse JWT Auth:")
		return
	}
	var profile profile
	row := db.DB.QueryRow(getMyProfile, email)
	err = row.Scan(&profile.FirstName, &profile.LastName, &profile.Email, &profile.PhoneNumber, &profile.Gender, &profile.BirthDate, &profile.IsVerified)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[GET PROFILE][ERROR] SELECT PROFILE:")
		return
	}
	middleware.ReturnResponseWriter(nil, w, getProfileResponse{
		Message: "Get Profile Success",
		Profile: profile,
	}, "[GET PROFILE][SUCCESS]")
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("[EDIT PROFILE][REQUEST]")
	email, err := middleware.ParseAuth(r.Header, middleware.KeyClient)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[EDIT PROFILE][ERROR] Parse JWT Auth:")
		return
	}
	var req editProfileRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[EDIT PROFILE][ERROR] DECODE REQUEST:")
		return
	}

	_, err = db.DB.Exec(editMyProfile, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Gender, req.BirthDate, email)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[EDIT PROFILE][ERROR] UPDATE DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, editProfileResponse{Message: "Success edit profile."}, "[EDIT PROFILE][SUCCESS]")
}
