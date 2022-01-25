package account

const (
	insertAccount      = `INSERT INTO accounts (first_name, last_name, email, password, phone_number, gender, birth_date, is_verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, false, now(), now())`
	updateVerification = `UPDATE accounts SET is_verified = $1, updated_at = now() WHERE email = $2`
	getPasswordUser    = `SELECT password FROM accounts WHERE email = $1`
)

type (
	registerRequest struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		Password    string `json:"password"`
		PhoneNumber string `json:"phone_number"`
		Gender      string `json:"gender"`
		BirthDate   string `json:"birthdate"`
	}

	registerResponse struct {
		Message string `json:"message"`
	}

	otpRequest struct {
		OTP   string `json:"otp"`
		Email string `json:"email"`
	}

	otpResponse struct {
		Message string `json:"message"`
	}

	resendOtpRequest struct {
		Email string `json:"email"`
	}

	resendOtpResponse struct {
		Message string `json:"message"`
	}

	loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginResponse struct {
		Message        string `json:"message"`
		Authentication string `json:"authentication"`
	}
)
