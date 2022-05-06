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
	"thegrace/pkg/services/booking"
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
	r := handleRouter()
	r.Use(middleware.DefaultHeader)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func handleRouter() *mux.Router {
	r := mux.NewRouter()
	r = r.PathPrefix("/api/v1").Subrouter()
	r.Use(middleware.DefaultHeader)
	accountRouter := r.PathPrefix("/user").Subrouter()
	accountRouter.HandleFunc("/register", account.RegisterAccountHandler).Methods("POST", "OPTIONS")
	accountRouter.HandleFunc("/login", account.LoginHandler).Methods("POST", "OPTIONS")
	accountRouter.HandleFunc("/profile", middleware.IsAuthorizedUser(profile.GetProfile)).Methods("GET", "OPTIONS")
	accountRouter.HandleFunc("/profile/edit", middleware.IsAuthorizedUser(profile.EditProfile)).Methods("POST", "OPTIONS")

	ibadahRouter := r.PathPrefix("/ibadah").Subrouter()
	ibadahRouter.HandleFunc("/khotbah/latest", ibadah.GetLatestKhotbah).Methods("GET", "OPTIONS")
	ibadahRouter.HandleFunc("/khotbah/list", ibadah.GetListKhotbah).Methods("GET", "OPTIONS")
	ibadahRouter.HandleFunc("/khotbah/jadwal", ibadah.GetJadwalIbadahById).Methods("GET", "OPTIONS")
	ibadahRouter.HandleFunc("/khotbah/jadwal/list", ibadah.GetJadwalIbadahList).Methods("GET", "OPTIONS")

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.HandleFunc("/login", admin.LoginAdmin).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/account/list", middleware.IsAuthorizedAdmin(admin.GetAccountList)).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/account/edit", middleware.IsAuthorizedAdmin(admin.EditProfile)).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/khotbah", middleware.IsAuthorizedAdmin(ibadah.AddNewKhotbah)).Methods("POST", "OPTIONS")
	adminRouter.HandleFunc("/khotbah/delete", middleware.IsAuthorizedAdmin(ibadah.DeleteKhotbahById)).Methods("GET", "OPTIONS")
	adminRouter.HandleFunc("/khotbah/jadwal/add", middleware.IsAuthorizedAdmin(ibadah.AddNewJadwalIbadah)).Methods("POST", "OPTIONS")

	bookingRouter := r.PathPrefix("/reservasi").Subrouter()
	bookingRouter.HandleFunc("/book", middleware.IsAuthorizedUser(booking.AddNewBooking)).Methods("POST", "OPTIONS")
	bookingRouter.HandleFunc("/book/user", middleware.IsAuthorizedUser(booking.GetBookingsFromBookerId)).Methods("GET", "OPTIONS")
	bookingRouter.HandleFunc("/book/usher", middleware.IsAuthorizedUsher(booking.GetBookingsFromIbadahId)).Methods("POST", "OPTIONS")
	bookingRouter.HandleFunc("/scan", middleware.IsAuthorizedUsher(booking.ScanReservation)).Methods("POST", "OPTIONS")
	bookingRouter.HandleFunc("/change_status", middleware.IsAuthorizedUsher(booking.ChangeBookingStatus)).Methods("POST", "OPTIONS")

	return r
}
