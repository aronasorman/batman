package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

var db gorm.DB

func main() {
	r := pat.New()
	r.Get("/alerts", getAlerts)
	r.Post("/alerts", newAlert)

	http.Handle("/", r)

	go func() {
		fmt.Println("Listening at port 8000")
	}()

	// initialize the db
	var err error
	db, err = initDb()
	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8000", nil)
}

func getAlerts(w http.ResponseWriter, r *http.Request) {
	// dummy alert, for testing
	alert := Alert{
		Id:       1,
		Location: "Tondo",
		Type:     "Murder",
	}

	// set application type to json
	w.Header().Set("Content-Type", "application/json")

	// encode the alerts to json, then write it to the client
	json.NewEncoder(w).Encode(alert)
}

func newAlert(w http.ResponseWriter, r *http.Request) {
	var alert Alert
	body := r.Body

	err := json.NewDecoder(body).Decode(&alert)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	if db.NewRecord(&alert) {
		db.Create(&alert)
	}

	w.Write([]byte("OK"))
}

func initDb() (db gorm.DB, err error) {
	// initialize the db
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		return
	}

	// setup and migrate the tables, if necessary.
	// AutoMigrate has some caveats. See documentation for
	// details.
	db.AutoMigrate(&Alert{})

	return
}
