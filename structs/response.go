package structs

type ResJabatan struct {
	IDJabatan   int    `form:"id" json:"id"`
	NamaJabatan string `form:"nama" json:"nama"`
}

type ResPegawai struct {
	IDPegawai int    `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	Alamat    string `form:"alamat" json:"alamat"`
}
