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
	"thegrace/pkg/services/ibadah"
	"thegrace/pkg/services/profile"
)

func init() {
	err := pkg.GetConfigEnv()
	if err != nil {
		panic(err)
	}
	err = db.GetConnection(pkg.Conf.DbHost, pkg.Conf.DbPort, pkg.Conf.DbUsername, pkg.Conf.DbPassword, pkg.Conf.DbName)
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Running Web Service. . .")
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()
	r = handler(r)
	r.Use(middleware.DefaultHeader)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Default port if not specified
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handler(r *mux.Router) *mux.Router {
	accountRouter := r.PathPrefix("/user").Subrouter()
	accountRouter.HandleFunc("/register", account.RegisterAccountHandler).Methods("POST")
	accountRouter.HandleFunc("/login", account.LoginHandler).Methods("POST")
	//accountRouter.HandleFunc("/otp", account.ValidateOTPHandler).Methods("POST")
	//accountRouter.HandleFunc("/otp/resend", account.ResendOTPHandler).Methods("POST")
	accountRouter.HandleFunc("/profile", middleware.IsAuthorizedUser(profile.GetProfile)).Methods("GET")
	accountRouter.HandleFunc("/profile/edit", middleware.IsAuthorizedUser(profile.EditProfile)).Methods("POST")

	ibadahRouter := r.PathPrefix("/ibadah").Subrouter()
	ibadahRouter.HandleFunc("/khotbah/latest", ibadah.GetLatestIbadah).Methods("GET")
	ibadahRouter.HandleFunc("/khotbah/list", ibadah.GetListKhotbah).Methods("GET")

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", admin.LoginAdmin).Methods("POST")
	adminRouter.HandleFunc("/account/list", middleware.IsAuthorizedAdmin(admin.GetAccountList)).Methods("GET")
	adminRouter.HandleFunc("/account/edit", middleware.IsAuthorizedAdmin(admin.EditProfile)).Methods("POST")
	adminRouter.HandleFunc("/khotbah", middleware.IsAuthorizedAdmin(ibadah.AddNewKhotbah)).Methods("POST")

	return r
}
