package profile

import "time"

const (
	GetMyProfile   = `SELECT id, first_name, last_name, email, phone_number, gender, birth_date, is_verified FROM accounts where email = $1`
	GetMyProfileId = `SELECT id FROM accounts WHERE email = $1`
	editMyProfile  = `UPDATE accounts SET first_name = $1, last_name = $2, email = $3, phone_number = $4, gender = $5, birth_date = $6, updated_at = now() WHERE email = $7`
)

type (
	getProfileResponse struct {
		Message string  `json:"message"`
		Profile Profile `json:"profile"`
	}

	Profile struct {
		Id          int       `json:"id"`
		FirstName   string    `field:"first_name" json:"first_name"`
		LastName    string    `field:"last_name" json:"last_name"`
		Email       string    `field:"email" json:"email"`
		PhoneNumber string    `field:"phone_number" json:"phone_number"`
		Gender      string    `field:"gender" json:"gender"`
		BirthDate   time.Time `field:"birth_date" json:"birth_date"`
		IsVerified  bool      `field:"is_verified" json:"is_verified"`
	}

	editProfileRequest struct {
		FirstName   string    `field:"first_name" json:"first_name"`
		LastName    string    `field:"last_name" json:"last_name"`
		Email       string    `field:"email" json:"email"`
		PhoneNumber string    `field:"phone_number" json:"phone_number"`
		Gender      string    `field:"gender" json:"gender"`
		BirthDate   time.Time `field:"birth_date" json:"birth_date"`
	}

	editProfileResponse struct {
		Message string `json:"message"`
	}
)
