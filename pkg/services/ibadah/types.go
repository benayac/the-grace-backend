package ibadah

const (
	insertKhotbah         = `INSERT INTO khotbah (thumbnail, title, link, pendeta_name, ibadah_date, link_warta) VALUES ($1, $2, $3, $4, $5, $6)`
	getKhotbahLatest      = `SELECT * FROM khotbah WHERE now() >= khotbah.ibadah_date ORDER BY khotbah.ibadah_date DESC LIMIT 1`
	getKhotbahListLimited = `SELECT * FROM KHOTBAH LIMIT $1`
	getKhotbahList        = `SELECT * FROM KHOTBAH`
	deleteKhotbahById     = `DELETE FROM khotbah WHERE id = $1`
	insertIbadah          = `INSERT INTO Ibadah (ibadah_title, location, date, max_capacity) VALUES ($1, $2, $3, $4)`
	GetIbadahById         = `SELECT * FROM Ibadah WHERE id = $1`
	GetIbadahList         = `SELECT * FROM Ibadah WHERE date >= now()`
	UpdateIbadahFilled    = `UPDATE ibadah set filled_capacity = $1 WHERE id = $2`
)

type (
	khotbah struct {
		Id          int    `json:"id"`
		Thumbnail   string `json:"thumbnail"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		PendetaName string `json:"pendeta_name"`
		IbadahDate  string `json:"ibadah_date"`
		LinkWarta   string `json:"link_warta"`
	}

	Ibadah struct {
		Id             int    `json:"id"`
		Title          string `json:"ibadah_title"`
		Location       string `json:"location"`
		IbadahDate     string `json:"ibadah_date"`
		MaxCapacity    int    `json:"max_capacity"`
		FilledCapacity int    `json:"filled_capacity"`
	}

	addKhotbahRequest struct {
		Thumbnail   string `json:"thumbnail"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		PendetaName string `json:"pendeta_name"`
		IbadahDate  string `json:"ibadah_date"`
		LinkWarta   string `json:"link_warta"`
	}

	addKhotbahResponse struct {
		Message string `json:"message"`
	}

	getKhotbahListResponse struct {
		Message string    `json:"message"`
		Khotbah []khotbah `json:"khotbah_list"`
	}

	getKhotbahLatestResponse struct {
		Message string  `json:"message"`
		Khotbah khotbah `json:"khotbah"`
	}

	addIbadahRequest struct {
		Title       string `json:"ibadah_title"`
		Location    string `json:"location"`
		IbadahDate  string `json:"ibadah_date"`
		MaxCapacity int    `json:"max_capacity"`
	}

	getIbadahByIdRequest struct {
		IbadahId int `json:"ibadah_id"`
	}

	getIbadahResponse struct {
		Message string `json:"message"`
		Ibadah  Ibadah `json:"Ibadah"`
	}

	getIbadahListResponse struct {
		Message string   `json:"message"`
		Ibadah  []Ibadah `json:"ibadah_list"`
	}
)
