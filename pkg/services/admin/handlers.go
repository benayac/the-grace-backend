package admin

import (
	"encoding/json"
	"log"
	"net/http"
	"thegrace/pkg"
	"thegrace/pkg/db"
	"thegrace/pkg/helper"
	"thegrace/pkg/middleware"
)

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	log.Println("[LOGIN ADMIN][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var login adminLoginRequest
	err := decoder.Decode(&login)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN ADMIN][ERROR] DECODE REQUEST:")
		return
	}
	var hash string

	defer db.DB.Close()
	err = db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN ADMIN][ERROR] CONNECT TO DB:")
		return
	}
	row := db.DB.QueryRow(getPasswordAdmin, login.Username)
	err = row.Scan(&hash)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN ADMIN][ERROR] SELECT PASSWORD:")
		return
	}
	validation, err := helper.CompareHashAndPassword(hash, []byte(login.Password))
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN ADMIN][ERROR] COMPARE HASH PASSWORD:")
		return
	}

	if validation {
		jwt, err := middleware.GetJWTAdmin(login.Username)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, nil, "[LOGIN ADMIN][ERROR] GET JWT AUTH:")
			return
		}
		middleware.ReturnResponseWriter(nil, w, adminLoginResponse{Message: "Login Success", Authentication: jwt}, "[LOGIN ADMIN][SUCCESS]")
	} else {
		middleware.ReturnResponseWriter(nil, w, adminLoginResponse{Message: "Credential Invalid"}, "[LOGIN ADMIN][SUCCESS]")
	}
}

func GetAccountList(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADMIN GET ACCOUNT LIST][REQUEST]")

	defer db.DB.Close()
	err := db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getAccountListResponse{Message: "Failed to get account list"}, "[ADMIN GET ACCOUNT LIST][ERROR] CONNECT TO DB:")
		return
	}
	row, err := db.DB.Query(getAccountList)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getAccountListResponse{Message: "Failed to get account list"}, "[ADMIN GET ACCOUNT LIST][ERROR] QUERY DB")
		return
	}

	var acc userProfile
	var accountList []userProfile
	for row.Next() {
		err = row.Scan(&acc.AccountId, &acc.FirstName, &acc.LastName, &acc.Email, &acc.PhoneNumber, &acc.Gender, &acc.BirthDate, &acc.IsVerified, &acc.Tag)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, getAccountListResponse{Message: "Failed to get account list"}, "[ADMIN GET ACCOUNT LIST][ERROR] QUERY ROW DB ")
			return
		}
		accountList = append(accountList, acc)
	}
	middleware.ReturnResponseWriter(nil, w, getAccountListResponse{Message: "Success to get account list", AccountList: accountList}, "[ADMIN GET ACCOUNT LIST][SUCCESS]")
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADMIN EDIT PROFILE][REQUEST]")
	var req userProfile
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[ADMIN EDIT PROFILE][ERROR] DECODE REQUEST:")
		return
	}

	defer db.DB.Close()
	err = db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[ADMIN EDIT PROFILE][ERROR] CONNECTION TO DB:")
		return
	}
	_, err = db.DB.Exec(editUserProfile, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Gender, req.BirthDate, req.IsVerified, req.Tag, req.AccountId)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[ADMIN EDIT PROFILE][ERROR] UPDATE DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, editProfileResponse{Message: "Success edit profile."}, "[ADMIN EDIT PROFILE][SUCCESS]")
}
