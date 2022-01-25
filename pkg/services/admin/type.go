package admin

const (
	getPasswordAdmin = `SELECT password FROM admins WHERE username = $1`
)

type (
	adminLoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	adminLoginResponse struct {
		Message        string `json:"message"`
		Authentication string `json:"authentication"`
	}
)
