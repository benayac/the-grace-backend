package admin

import (
	"time"
)

const (
	getPasswordAdmin  = `SELECT password FROM admins WHERE username = $1`
	getAccountList    = `SELECT id, first_name, last_name, email, phone_number, gender, birth_date, is_verified, tag FROM accounts`
	editUserProfile   = `UPDATE accounts SET first_name = $1, last_name = $2, email = $3, phone_number = $4, gender = $5, birth_date = $6, is_verified = $7, tag = $8, updated_at = now() WHERE id = $9`
	deleteUserProfile = `DELETE FROM accounts WHERE id = $1`
)

type (
	userProfile struct {
		AccountId   int       `json:"id"`
		FirstName   string    `field:"first_name" json:"first_name"`
		LastName    string    `field:"last_name" json:"last_name"`
		Email       string    `field:"email" json:"email"`
		PhoneNumber string    `field:"phone_number" json:"phone_number"`
		Gender      string    `field:"gender" json:"gender"`
		BirthDate   time.Time `field:"birth_date" json:"birth_date"`
		IsVerified  bool      `field:"is_verified" json:"is_verified"`
		Tag         int       `json:"tag"`
	}

	adminLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	adminLoginResponse struct {
		Message        string `json:"message"`
		Authentication string `json:"authentication"`
	}

	getAccountListResponse struct {
		Message     string        `json:"message"`
		AccountList []userProfile `json:"account_list"`
	}

	editProfileResponse struct {
		Message string `json:"message"`
	}
)
