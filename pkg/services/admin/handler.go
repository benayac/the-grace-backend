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
