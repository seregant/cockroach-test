package structs

type ResJabatan struct {
	IDJabatan   string `form:"id" json:"id"`
	NamaJabatan string `form:"nama" json:"nama"`
}

type ResDivisi struct {
	IDDivisi   string `form:"id" json:"id"`
	NamaDivisi string `form:"nama" json:"nama"`
}

type ResPegawai struct {
	IDPegawai string `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	Alamat    string `form:"alamat" json:"alamat"`
	Username  string `form:"username" json:"username"`
	Password  string `form:"password" json:"password"`
	Email     string `form:"email" json:"email"`
	JabatanID string `form:"jabatan" json:"jabatan"`
	DivisiID  string `form:"divisi" json:"divisi"`
}

type ResPekerjaan struct {
	IDPekerjaan   int      `form:"id" json:"id"`
	NamaPekerjaan string   `form:"nama" json:"nama"`
	NamaPj        string   `form:"pj" json:"pj"`
	Tim           []string `form:"anggota_tim" json:"anggota_tim"`
	Deadline      string   `form:"deadline" json:"deadline"`
}
