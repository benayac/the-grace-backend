package ibadah

import (
	"database/sql"
	"log"
	"net/http"
	"thegrace/pkg/db"
	"thegrace/pkg/middleware"
)

func GetLatestIbadah(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET LATEST KHOTBAH][REQUEST]")

	var khotbah khotbah
	row := db.DB.QueryRow(getKhotbahLatest)
	err := row.Scan(&khotbah.Id, &khotbah.Thumbnail, &khotbah.Title, &khotbah.Link, &khotbah.PendetaName, &khotbah.IbadahDate, &khotbah.LinkWarta)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getKhotbahListResponse{Message: "Failed to get khotbah latest"}, "[GET LATEST KHOTBAH][ERROR] QUERY DATA DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, getKhotbahLatestResponse{Message: "Success to get khotbah latest", Khotbah: khotbah}, "[GET LATEST KHOTBAH][SUCCESS]")
}

func GetListKhotbah(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET LIST KHOTBAH][REQUEST]")
	limit := r.FormValue("limit")
	var row *sql.Rows
	var err error
	if limit != "" {
		row, err = db.DB.Query(getKhotbahListLimited, limit)
	} else {
		row, err = db.DB.Query(getKhotbahList)
	}
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getKhotbahListResponse{Message: "Failed to get khotbah list"}, "[GET LIST KHOTBAH][ERROR] QUERY DATA DB:")
		return
	}
	var list []khotbah
	var khotbah khotbah
	for row.Next() {
		err = row.Scan(&khotbah.Id, &khotbah.Thumbnail, &khotbah.Title, &khotbah.Link, &khotbah.PendetaName, &khotbah.IbadahDate, &khotbah.LinkWarta)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, getKhotbahListResponse{Message: "Failed to get khotbah list"}, "[GET LIST KHOTBAH][ERROR] QUERY ROW DB:")
			return
		}
		list = append(list, khotbah)
	}
	middleware.ReturnResponseWriter(nil, w, getKhotbahListResponse{Message: "Success to get khotbah list", Khotbah: list}, "[GET LIST KHOTBAH][SUCCESS]")
}
