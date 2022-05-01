package account

import (
	"encoding/json"
	"log"
	"net/http"
	"thegrace/pkg/db"
	"thegrace/pkg/helper"
	"thegrace/pkg/middleware"
)

func RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[REGISTER ACCOUNT][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req registerRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, registerResponse{Message: "Failed to register"}, "[REGISTER ACCOUNT][ERROR] DECODE REQUEST:")
	}

	hashedPass, err := helper.HashAndSalt([]byte(req.Password))
	if err != nil {
		middleware.ReturnResponseWriter(err, w, registerResponse{Message: "Failed to register"}, "[REGISTER ACCOUNT][ERROR] FAIL TO HASH PASSWORD:")
		return
	}
	req.Password = hashedPass

	_, err = db.DB.Exec(insertAccount, req.FirstName, req.LastName, req.Email, req.Password,
		req.PhoneNumber, req.Gender, req.BirthDate)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, registerResponse{Message: "Failed to register"}, "[REGISTER ACCOUNT][ERROR] INSERT TO DB:")
		return
	}
	//otp, err := helper.GenerateOTP(req.Email)
	if err != nil {
		log.Println("[REGISTER ACCOUNT][ERROR] GENERATE OTP:" + err.Error())
		return
	}
	//go func() {
	//	err = helper.SendEmail(pkg.Conf.Email, pkg.Conf.EmailPassword, req.Email, "OTP", otp)
	//	if err != nil {
	//		log.Println("[REGISTER ACCOUNT][ERROR] SENDING EMAIL:" + err.Error())
	//	}
	//	log.Println("[REGISTER ACCOUNT][EMAIL SENT]")
	//}()
	middleware.ReturnResponseWriter(nil, w, registerResponse{Message: "Register Success"}, "[REGISTER ACCOUNT][SUCCESS]")
}

/**
func ResendOTPHandler(w http.ResponseWriter, r *http.Request) {
	var req resendOtpRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		log.Println("[RESEND OTP][ERROR] DECODE REQUEST:" + err.Error())
		return
	}

	otp, err := helper.GenerateOTP(req.Email)
	if err != nil {
		log.Println("[RESEND OTP][ERROR] GENERATE OTP:" + err.Error())
		return
	}
	go func() {
		err = helper.SendEmail(pkg.Conf.Email, pkg.Conf.EmailPassword, req.Email, "OTP", otp)
		if err != nil {
			log.Println("[REGISTER ACCOUNT][ERROR] SENDING EMAIL:" + err.Error())
		}
		log.Println("[REGISTER ACCOUNT][EMAIL SENT]")
	}()
	middleware.ReturnResponseWriter(nil, w, resendOtpResponse{Message: "Success to resend OTP"}, "[RESEND OTP][SUCCESS]")
}

func ValidateOTPHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[VALIDATE OTP][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var otp otpRequest
	err := decoder.Decode(&otp)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, otpResponse{Message: "Failed to verify OTP"}, "[VALIDATE OTP][ERROR] DECODE REQUEST:")
		return
	}

	result, err := helper.ValidateOTP(otp.OTP, otp.Email)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, otpResponse{Message: "Failed to verify OTP"}, "[VALIDATE OTP][ERROR] VALIDATE OTP:")
		return
	}
	if result {
		_, err = db.DB.Exec(updateVerification, true, otp.Email)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, otpResponse{Message: "Failed to verify OTP"}, "[VALIDATE OTP][ERROR] INSERT TO DB:")
			return
		}
	} else {
		middleware.ReturnResponseWriter(nil, w, otpResponse{Message: "Failed to verify OTP"}, "[VALIDATE OTP][ERROR] AUTH INVALID:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, otpResponse{Message: "Verification Success"}, "[VALIDATE OTP][SUCCESS]")
}

*/

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[LOGIN][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var login loginRequest
	err := decoder.Decode(&login)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN][ERROR] DECODE REQUEST:")
		return
	}
	var hash string
	var tag int
	row := db.DB.QueryRow(getPasswordUser, login.Email)
	err = row.Scan(&hash, &tag)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN][ERROR] SELECT PASSWORD:")
		return
	}
	validation, err := helper.CompareHashAndPassword(hash, []byte(login.Password))
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[LOGIN][ERROR] COMPARE HASH PASSWORD:")
		return
	}

	if validation {
		jwt, err := middleware.GetJWTUser(login.Email)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, loginResponse{Message: "Login Failed"}, "[LOGIN][ERROR] GET JWT USER AUTH: "+err.Error())
			return
		}
		if tag == accountTypeUser {
			middleware.ReturnResponseWriter(nil, w, loginResponse{Message: "Login Success", AccountType: tag, Authentication: authenticationStruct{UserJwt: jwt}}, "[LOGIN][SUCCESS]")
		} else if tag == accountTypeUsher {
			usherJwt, err := middleware.GetJWTUsher(login.Email)
			if err != nil {
				middleware.ReturnResponseWriter(err, w, loginResponse{Message: "Login Failed"}, "[LOGIN][ERROR] GET JWT USHER AUTH: "+err.Error())
				return
			}
			middleware.ReturnResponseWriter(nil, w, loginResponse{Message: "Login Success", AccountType: tag, Authentication: authenticationStruct{UserJwt: jwt, UsherJwt: usherJwt}}, "[LOGIN][SUCCESS]")
		}
	} else {
		middleware.ReturnResponseWriter(nil, w, loginResponse{Message: "Credential Invalid"}, "[LOGIN][SUCCESS]")
	}
}
