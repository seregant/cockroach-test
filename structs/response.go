package structs

type ResJabatan struct {
	IDJabatan   int    `form:"id" json:"id"`
	NamaJabatan string `form:"nama" json:"nama"`
}

type DataResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []ResJabatan
}

type StdrResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
