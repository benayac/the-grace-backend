package ibadah

const (
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

	getKhotbahListResponse struct {
		Message string    `json:"message"`
		Khotbah []khotbah `json:"khotbah_list"`
	}

	getKhotbahLatestResponse struct {
		Message string  `json:"message"`
		Khotbah khotbah `json:"khotbah"`
	}
)
