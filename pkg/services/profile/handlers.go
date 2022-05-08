package profile

import (
	"encoding/json"
	"log"
	"net/http"
	"thegrace/pkg"
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
	var profile Profile

	defer db.DB.Close()
	err = db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[GET PROFILE][ERROR] FAILED CONNECTION TO DB:")
		return
	}

	row := db.DB.QueryRow(GetMyProfile, email)
	err = row.Scan(&profile.Id, &profile.FirstName, &profile.LastName, &profile.Email, &profile.PhoneNumber, &profile.Gender, &profile.BirthDate, &profile.IsVerified)
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
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit Profile."}, "[EDIT PROFILE][ERROR] Parse JWT Auth:")
		return
	}
	var req editProfileRequest
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit Profile."}, "[EDIT PROFILE][ERROR] DECODE REQUEST:")
		return
	}

	defer db.DB.Close()
	err = db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit Profile."}, "[EDIT PROFILE][ERROR] FAILED CONNECTION TO DB:")
		return
	}

	_, err = db.DB.Exec(editMyProfile, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Gender, req.BirthDate, email)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit Profile."}, "[EDIT PROFILE][ERROR] UPDATE DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, editProfileResponse{Message: "Success edit Profile."}, "[EDIT PROFILE][SUCCESS]")
}
