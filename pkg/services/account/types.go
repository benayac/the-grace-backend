package account

const (
	insertAccount = `INSERT INTO accounts (first_name, last_name, email, password, phone_number, gender, birth_date, is_verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, false, now(), now())`
	//updateVerification = `UPDATE accounts SET is_verified = $1, updated_at = now() WHERE email = $2`
	getPasswordUser  = `SELECT password, tag FROM accounts WHERE email = $1`
	accountTypeUser  = 1
	accountTypeUsher = 2
)

type (
	registerRequest struct {
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		Password    string `json:"password,omitempty"`
		PhoneNumber string `json:"phone_number,omitempty"`
		Gender      string `json:"gender,omitempty"`
		BirthDate   string `json:"birthdate,omitempty"`
	}

	registerResponse struct {
		Message string `json:"message,omitempty"`
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

	authenticationStruct struct {
		UserJwt  string `json:"user_jwt,omitempty"`
		UsherJwt string `json:"usher_jwt,omitempty"`
	}

	loginResponse struct {
		Message        string               `json:"message,omitempty"`
		Authentication authenticationStruct `json:"authentication,omitempty"`
		AccountType    int                  `json:"account_type,omitempty"`
	}
)
