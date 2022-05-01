package booking

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"thegrace/pkg"
	"thegrace/pkg/db"
	"thegrace/pkg/helper"
	"thegrace/pkg/middleware"
	"thegrace/pkg/services/ibadah"
	"thegrace/pkg/services/profile"
)

func AddNewBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("[ADD BOOKING][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req createBookingRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] DECODE REQUEST: ")
		return
	}

	var ib ibadah.Ibadah
	row := db.DB.QueryRow(ibadah.GetIbadahById, req.IbadahId)
	err = row.Scan(&ib.Id, &ib.Title, &ib.Location, &ib.IbadahDate, &ib.MaxCapacity, &ib.FilledCapacity)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] QUERY DB: ")
		return
	}

	if len(req.BookingData)+ib.FilledCapacity > ib.MaxCapacity {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Maximum capacity reached."}, "[ADD BOOKING][CANCELED] NOT ENOUGH CAPACITY: ")
		return
	}

	for b := range req.BookingData {
		var id int
		booking := req.BookingData[b]
		err := db.DB.QueryRow(insertBooking, req.BookerId, req.BookerName, booking.Name, req.IbadahId).Scan(&id)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] INSERT TO DB: ")
			return
		}
		data := fmt.Sprintf("id:%v,nama:%v,phoneNumber:%v,id_ibadah:%v", id, booking.Name, booking.PhoneNumber, req.IbadahId)
		encryptedData, err := helper.EncryptData(data, pkg.Conf.SigningKeyEncrypt)
		fmt.Printf("DATA: %s\n", encryptedData)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] ENCRYPT DATA: ")
			return
		}
		_, err = db.DB.Exec(addEncryptedData, encryptedData, id)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] UPDATE ENCRYPTED DATA: ")
			return
		}

	}
	_, err = db.DB.Exec(ibadah.UpdateIbadahFilled, ib.FilledCapacity+len(req.BookingData), ib.Id)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to create booking"}, "[ADD BOOKING][ERROR] INCREMENT FILLED CAPACITY DATA: ")
		return
	}
	middleware.ReturnResponseWriter(nil, w, createBookingResponse{Message: "Succes to create booking"}, "[ADD BOOKING][SUCCESS]")
}

func GetBookingsFromBookerId(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET BOOKINGS FROM BOOKER ID][REQUEST]")

	email, err := middleware.ParseAuth(r.Header, middleware.KeyClient)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getBookingResponse{Message: "Failed to get booking from booker id."}, "[GET BOOKINGS FROM BOOKER ID][ERROR] Parse JWT Auth:")
		return
	}
	var p profile.Profile
	row := db.DB.QueryRow(profile.GetMyProfileId, email)
	err = row.Scan(&p.Id)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getBookingResponse{Message: "Failed to get booking from booker id."}, "[GET BOOKINGS FROM BOOKER ID][ERROR] failed query profile db: ")
		return
	}

	rows, err := db.DB.Query(getBookingsDataFromBookerId, p.Id)
	var listBooking []bookingsData
	for rows.Next() {
		var booking bookingsData
		err = rows.Scan(&booking.Id, &booking.BookerId, &booking.BookerName, &booking.Name, &booking.BookingDate, &booking.IbadahId, &booking.EncryptedData, &booking.Status)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, getBookingResponse{Message: "Failed to get booking from booker id."}, "[GET BOOKINGS FROM BOOKER ID][ERROR] failed query booking db: ")
			return
		}
		listBooking = append(listBooking, booking)
	}
	middleware.ReturnResponseWriter(nil, w, getBookingResponse{Message: "Success to get booking from booker id.", BookingList: listBooking}, "[GET BOOKINGS FROM BOOKER ID][SUCCESS]")
}

func GetBookingsFromIbadahId(w http.ResponseWriter, r *http.Request) {
	log.Println("[GET BOOKINGS FROM IBADAH ID][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req getBookingFromIbadahIdRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, getBookingResponse{Message: "Failed to get booking from ibadah id."}, "[GET BOOKINGS FROM IBADAH ID][ERROR] failed decode request: ")
		return
	}

	rows, err := db.DB.Query(getBookingsDataFromIbadahId, req.IbadahId)
	var listBooking []bookingsData
	for rows.Next() {
		var booking bookingsData
		err = rows.Scan(&booking.Id, &booking.BookerId, &booking.BookerName, &booking.Name, &booking.IbadahId, &booking.Status)
		if err != nil {
			middleware.ReturnResponseWriter(err, w, getBookingResponse{Message: "Failed to get booking from ibadah id."}, "[GET BOOKINGS FROM IBADAH ID][ERROR] failed query booking db: ")
			return
		}
		listBooking = append(listBooking, booking)
	}
	middleware.ReturnResponseWriter(nil, w, getBookingResponse{Message: "Success to get booking from ibadah id.", BookingList: listBooking}, "[GET BOOKINGS FROM IBADAH ID][SUCCESS]")
}

func ScanReservation(w http.ResponseWriter, r *http.Request) {
	log.Println("[SCAN RESERVATION][REQUEST]")

	decoder := json.NewDecoder(r.Body)
	var req scanReservationRequest
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to scan reservation."}, "[SCAN RESERVATION][ERROR] FAILED TO DECODE: ")
		return
	}

	bookingData, err := decryptBookingData(req.EncryptedData)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to scan reservation."}, "[SCAN RESERVATION][ERROR] FAILED TO DECRYPT BOOKING DATA: ")
		return
	}

	_, err = db.DB.Exec(updateBookingStatusFromBookingId, bookingData.Id)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to scan reservation."}, "[SCAN RESERVATION][ERROR] FAILED TO UPDATE STATUS: ")
		return
	}

	middleware.ReturnResponseWriter(nil, w, createBookingResponse{Message: "Success to scan reservation."}, "[SCAN RESERVATION][SUCCESS]")
}

func ChangeBookingStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("[CHANGE BOOKING STATUS][REQUEST]")
	var req changeBookingStatusRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to change booking status."}, "[CHANGE BOOKING STATUS][ERROR] FAILED TO DECODE: ")
		return
	}

	_, err = db.DB.Exec(changeBookingStatusFromBookingId, req.Status, req.BookingId)
	if err != nil {
		middleware.ReturnResponseWriter(err, w, createBookingResponse{Message: "Failed to change booking status."}, "[CHANGE BOOKING STATUS][ERROR] FAILED TO EXECUTE QUERY: ")
		return
	}

	middleware.ReturnResponseWriter(nil, w, createBookingResponse{Message: "Success to change booking status."}, "[CHANGE BOOKING STATUS][SUCCESS]")
}
