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
	alert := Alert{
		Id:       1,
		Location: "Tondo",
		Type:     "Murder",
	}
	json.NewEncoder(w).Encode(alert)
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
