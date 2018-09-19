package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-Blog/db"

	"github.com/go-Blog/model"
	"github.com/gorilla/mux"
)

// ResponseWithJSON return json data
func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetName test get name handler
func GetName(w http.ResponseWriter, r *http.Request) {
	var godusers []model.Goduser
	goduser := model.Goduser{}

	stmt, _ := db.DB.Prepare(`SELECT * FROM  customer WHERE AGE <?`)
	defer stmt.Close()
	rows, err := stmt.Query(33)
	if err != nil {
		log.Printf("select data error: %v\n", err.Error())
	}
	for rows.Next() {
		rows.Scan(&goduser.CUST_ID, &goduser.NAME, &goduser.AGE)
		godusers = append(godusers, goduser)
	}
	ResponseWithJSON(w, http.StatusOK, model.Response{
		Code:    0,
		Message: "ok",
		Data:    godusers,
	})
}

// GetAge test get age handler
func GetAge(w http.ResponseWriter, r *http.Request) {
	age := mux.Vars(r)["age"]
	// w.Header().Set("Content-Type","application/json;charset=UTF-8")
	w.Write([]byte("age is " + age))
}
