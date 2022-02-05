package admin

import (
	"encoding/json"
	"log"
	"net/http"
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
			middleware.ReturnResponseWriter(err, w, getAccountListResponse{Message: "Failed to get account list"}, "[ADMIN GET ACCOUNT LIST][ERROR] QUERY ROW DB")
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

	_, err = db.DB.Exec(editUserProfile, req.FirstName, req.LastName, req.Email, req.PhoneNumber, req.Gender, req.BirthDate, req.IsVerified, req.Tag, req.AccountId)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit profile."}, "[ADMIN EDIT PROFILE][ERROR] UPDATE DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, editProfileResponse{Message: "Success edit profile."}, "[ADMIN EDIT PROFILE][SUCCESS]")
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADMIN DELETE ACCOUNT][REQUEST]")
	accountId := r.FormValue("id")

	if accountId == "" {
		middleware.ReturnResponseWriter(nil, w, deleteProfileResponse{Message: "Failed to delete account"}, "[ADMIN DELETE ACCOUNT][ERROR] NO ID")
		return
	}

	_, err := db.DB.Exec(deleteUserProfile, accountId)

	if err != nil {
		middleware.ReturnResponseWriter(err, w, deleteProfileResponse{Message: "Failed to delete account."}, "[ADMIN DELETE ACCOUNT][ERROR] EXECUTE DB")
		return
	}

	middleware.ReturnResponseWriter(nil, w, deleteProfileResponse{Message: "Success delete account."}, "[ADMIN DELETE ACCOUNT][SUCCESS]")
}

func AddNewKhotbah(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADD KHOTBAH][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req addKhotbahRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[ADD KHOTBAH][ERROR] DECODE REQUEST:")
		return
	}
	_, err = db.DB.Exec(insertKhotbah, req.Thumbnail, req.Title, req.Link, req.PendetaName, req.IbadahDate, req.LinkWarta)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, addKhotbahResponse{Message: "Failed to add khotbah"}, "[ADD KHOTBAH][ERROR] INSERT TO DB:")
		return
	}
	middleware.ReturnResponseWriter(nil, w, addKhotbahResponse{Message: "Success to add khotbah"}, "[ADD KHOTBAH][SUCCESS]")
}

func GetListKhotbah(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADMIN GET LIST KHOTBAH][REQUEST]")
	row, err := db.DB.Query(getKhotbahList)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getKhotbahListResponse{Message: "Failed to get khotbah list"}, "[ADMIN GET LIST KHOTBAH][ERROR] QUERY DATA DB:")
		return
	}
	var list []khotbah
	var khotbah khotbah
	for row.Next() {
		err = row.Scan(&khotbah.Id, &khotbah.Thumbnail, &khotbah.Title, &khotbah.Link, &khotbah.PendetaName, &khotbah.IbadahDate, &khotbah.LinkWarta)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, getKhotbahListResponse{Message: "Failed to get khotbah list"}, "[ADMIN GET LIST KHOTBAH][ERROR] QUERY ROW DB:")
			return
		}
		list = append(list, khotbah)
	}
	middleware.ReturnResponseWriter(nil, w, getKhotbahListResponse{Message: "Success to get khotbah list", Khotbah: list}, "[ADMIN GET LIST KHOTBAH][SUCCESS]")
}

func EditKhotbah(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADMIN EDIT KHOTBAH][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req khotbah
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit khotbah"}, "[ADMIN EDIT KHOTBAH][ERROR] DECODE REQUEST")
		return
	}

	_, err = db.DB.Exec(editKhotbah, req.Thumbnail, req.Title, req.Link, req.PendetaName, req.IbadahDate, req.LinkWarta, req.Id)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, editProfileResponse{Message: "Failed to edit khotbah"}, "[ADMIN EDIT KHOTBAH][ERROR] EXEC DB")
		return
	}

	middleware.ReturnResponseWriter(nil, w, editProfileResponse{Message: "Success to edit khotbah"}, "[ADMIN EDIT KHOTBAH][SUCCESS]")
}
