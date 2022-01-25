package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"thegrace/pkg"
	"thegrace/pkg/db"
	"thegrace/pkg/middleware"
	"thegrace/pkg/services/account"
	"thegrace/pkg/services/admin"
	"thegrace/pkg/services/profile"
)

func init() {
	profile := os.Getenv("ENVIRONMENT")
	if profile == "LOCAL" {
		err := pkg.GetConfigJson()
		if err != nil {
			panic(err)
		}
	} else if profile == "DOCKER" {
		err := pkg.GetConfigEnv()
		if err != nil {
			panic(err)
		}
	}
	err := db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Running Web Service. . .")
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()
	r = handleRouter(r)
	r.Use(middleware.DefaultHeader)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handleRouter(r *mux.Router) *mux.Router {
	accountRouter := r.PathPrefix("/user").Subrouter()
	accountRouter.HandleFunc("/register", account.RegisterAccountHandler).Methods("POST")
	accountRouter.HandleFunc("/login", account.LoginHandler).Methods("POST")
	accountRouter.HandleFunc("/otp", account.ValidateOTPHandler).Methods("POST")
	accountRouter.HandleFunc("/otp/resend", account.ResendOTPHandler).Methods("POST")
	accountRouter.HandleFunc("/profile", middleware.IsAuthorizedUser(profile.GetProfile)).Methods("GET")
	accountRouter.HandleFunc("/profile/edit", middleware.IsAuthorizedUser(profile.EditProfile)).Methods("POST")

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", admin.LoginAdmin).Methods("POST")
	return r
}
