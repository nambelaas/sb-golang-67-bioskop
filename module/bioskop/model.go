package bioskop

type Bioskop struct {
	Id     int     `json:"id"`
	Nama   string  `json:"nama" binding:"required"`
	Lokasi string  `json:"lokasi" binding:"required"`
	Rating float64 `json:"rating"`
}
