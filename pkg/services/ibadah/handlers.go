package ibadah

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"thegrace/pkg/db"
	"thegrace/pkg/middleware"
)

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
