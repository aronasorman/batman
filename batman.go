package main

import (
	"fmt"
	"github.com/gorilla/pat"
	"github.com/jinzhu/gorm"
	"net/http"
)

func main() {
	r := pat.New()
	r.Get("/products", hello)

	http.Handle("/", r)

	go func() {
		fmt.Println("Listening at port 8000")
	}()

	// initialize the db
	initDb()

	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

}

func initDb() (db gorm.DB, err error) {
	// initialize the db
	db, err = gorm.Open("sqlite3", "/tmp/gorm.db")

	// setup and migrate the tables, if necessary.
	// AutoMigrate has some caveats. See documentation for
	// details.
	db.CreateTable(&Alert{})
	db.AutoMigrate(&Alert{})

	return
}
