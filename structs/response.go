package structs

import "github.com/dgrijalva/jwt-go"

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

type Token struct {
	UserId uint
	jwt.StandardClaims
}
