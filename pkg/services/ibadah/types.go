package ibadah

const (
	insertKhotbah         = `INSERT INTO khotbah (thumbnail, title, link, pendeta_name, ibadah_date, link_warta) VALUES ($1, $2, $3, $4, $5, $6)`
	getKhotbahLatest      = `SELECT * FROM khotbah WHERE now() >= khotbah.ibadah_date ORDER BY khotbah.ibadah_date DESC LIMIT 1`
	getKhotbahListLimited = `SELECT * FROM KHOTBAH LIMIT $1`
	getKhotbahList        = `SELECT * FROM KHOTBAH`
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
)
