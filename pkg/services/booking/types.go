package booking

const (
	insertBooking                    = `INSERT INTO bookings (id_pemesan, name_pemesan, name, booking_date, id_ibadah, status) VALUES ($1, $2, $3, now(), $4, false) RETURNING id`
	addEncryptedData                 = `UPDATE bookings SET encrypted_data = $1 WHERE id = $2`
	getBookingsDataFromBookerId      = `SELECT * FROM bookings WHERE id_pemesan = $1`
	getBookingsDataFromIbadahId      = `SELECT id, id_pemesan, name_pemesan, name, id_ibadah, status FROM bookings WHERE id_ibadah = $1`
	updateBookingStatusFromBookingId = `UPDATE bookings SET status = true WHERE id = $1`
	changeBookingStatusFromBookingId = `UPDATE bookings SET status = $1 WHERE id = $2`
)

type bookingEntity struct {
	Name        string `json:"name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type bookingsData struct {
	Id            int    `json:"id,omitempty"`
	BookerId      int    `json:"booker_id,omitempty"`
	BookerName    string `json:"booker_name,omitempty"`
	Name          string `json:"name,omitempty"`
	BookingDate   string `json:"booking_date,omitempty"`
	IbadahId      int    `json:"ibadah_id,omitempty"`
	EncryptedData string `json:"encrypted_data,omitempty"`
	Status        bool   `json:"booking_status"`
}

type createBookingRequest struct {
	BookerId    int             `json:"booker_id,omitempty"`
	BookerName  string          `json:"booker_name,omitempty"`
	BookingData []bookingEntity `json:"booking_data,omitempty"`
	IbadahId    int             `json:"ibadah_id,omitempty"`
}

type createBookingResponse struct {
	Message string `json:"message,omitempty"`
}

type scanBookingResponse struct {
	Message string `json:"message,omitempty"`
	Name    string `json:"name,omitempty"`
}

type getBookingResponse struct {
	Message     string         `json:"message,omitempty"`
	BookingList []bookingsData `json:"booking_list,omitempty"`
}

type getBookingFromIbadahIdRequest struct {
	IbadahId int `json:"ibadah_id,omitempty"`
}

type scanReservationRequest struct {
	EncryptedData string `json:"encrypted_data,omitempty"`
}

type changeBookingStatusRequest struct {
	BookingId int  `json:"booking_id,omitempty"`
	Status    bool `json:"status,omitempty"`
}
