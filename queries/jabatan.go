package queries

import (
	"cockroach-test/database"
	"cockroach-test/structs"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllJabatan(w http.ResponseWriter, r *http.Request) {

	var arr_jabatan []structs.ResJabatan
	var response structs.DataResponse
	var jabatan []structs.Jabatan

	var db = database.DbConnect()
	defer db.Close()

	db.Find(&jabatan)

	for _, jabat := range jabatan {
		arr_jabatan = append(arr_jabatan, structs.ResJabatan{IDJabatan: jabat.IDJabatan, NamaJabatan: jabat.NamaJabatan})
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arr_jabatan

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func UpdateJabatan(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	} else {
		var response structs.StdrResponse
		r.ParseForm()

		urlVars := mux.Vars(r)
		var IDtoEdit = urlVars["id"]

		db := database.DbConnect()
		db.LogMode(true)
		defer db.Close()

		update := db.Model(&structs.Jabatan{}).Where("jabatan_id = ?", IDtoEdit).Update("NamaJabatan", r.FormValue("NamaJabatan"))

		if update.GetErrors() != nil {

			type errMsg struct {
				message string
			}

			var errData []errMsg

			for _, err := range update.GetErrors() {
				errData = append(errData, errMsg{message: err.Error()})
			}

			response.Status = 204
			response.Message = errData[0].message

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		} else {
			response.Status = 200
			response.Message = "Success"

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	}
}
