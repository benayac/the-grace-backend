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

func GetLatestKhotbah(w http.ResponseWriter, r *http.Request) {
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

func DeleteKhotbahById(w http.ResponseWriter, r *http.Request) {
	log.Println("[DELETE KHOTBAH BY ID][REQUEST]")
	id := r.FormValue("id")
	if id == "" {
		middleware.ReturnResponseWriter(nil, w, addKhotbahResponse{Message: "No Id Param Found."}, "[DELETE KHOTBAH BY ID][ERROR] NO ID PARAM FOUND ")
		return
	}

	res, err := db.DB.Exec(deleteKhotbahById, id)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, addKhotbahResponse{Message: "Error delete khotbah."}, "[DELETE KHOTBAH BY ID][ERROR] QUERY DB")
		return
	}

	count, err := res.RowsAffected()
	if err != nil {
		middleware.ReturnResponseWriter(err, w, addKhotbahResponse{Message: "Error delete khotbah."}, "[DELETE KHOTBAH BY ID][ERROR] CHECK COUNT")
		return
	}

	if count <= 0 {
		middleware.ReturnResponseWriter(nil, w, addKhotbahResponse{Message: "No khotbah effected"}, "[DELETE KHOTBAH BY ID][ERROR] NO KHOTBAH EFFECTED")
		return
	}

	middleware.ReturnResponseWriter(nil, w, addKhotbahResponse{Message: "Khotbah " + id + " deleted succesfully"}, "[DELETE KHOTBAH BY ID][SUCCESS]")
}

func AddNewJadwalIbadah(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADD IBADAH][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req addIbadahRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, nil, "[ADD IBADAH][ERROR] DECODE REQUEST:")
		return
	}
	_, err = db.DB.Exec(insertIbadah, req.Title, req.Location, req.IbadahDate, req.MaxCapacity)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, addKhotbahResponse{Message: "Failed to add Ibadah"}, "[ADD IBADAH][ERROR] INSERT TO DB:")
		return
	}
	middleware.ReturnResponseWriter(nil, w, addKhotbahResponse{Message: "Success to add Ibadah"}, "[ADD IBADAH][SUCCESS]")
}

func GetJadwalIbadahById(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET IBADAH BY ID][REQUEST]")
	id := r.FormValue("id")
	if id == "" {
		middleware.ReturnResponseWriter(nil, w, getIbadahResponse{Message: "No Id Param Found."}, "[GET IBADAH BY ID][ERROR] NO ID PARAM FOUND ")
		return
	}

	var ibadah Ibadah
	row := db.DB.QueryRow(GetIbadahById, id)
	err := row.Scan(&ibadah.Id, &ibadah.Title, &ibadah.Location, &ibadah.IbadahDate, &ibadah.MaxCapacity, &ibadah.FilledCapacity)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getIbadahResponse{Message: "Failed to get Ibadah latest"}, "[GET IBADAH BY ID][ERROR] QUERY DATA DB:")
		return
	}

	middleware.ReturnResponseWriter(nil, w, getIbadahResponse{Message: "Success to get Ibadah latest", Ibadah: ibadah}, "[GET IBADAH][SUCCESS]")
}
