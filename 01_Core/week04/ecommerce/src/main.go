package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type sauce struct {
	id      int
	name    string
	brand   string
	price   float64
	spice   int
	size    float64
	imgPath string
}

var db *sql.DB

func connectDB() *sql.DB {
	fmt.Println("Opening mysql connection")
	db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3307)/ecommerce")
	if err != nil {
		panic(err)
	}

	return db
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("Home page hit")
}

func getAllSauces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rows, err := db.Query("select * from sauces;")

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	sauces := make([]sauce, 0)
	for rows.Next() {
		var item sauce
		err = rows.Scan(&item.id, &item.name, &item.price, &item.spice, &item.size, &item.imgPath, &item.brand)
		if err != nil {
			panic(err)
		}
		sauces = append(sauces, item)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(sauces)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/sauces", getAllSauces)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {

	db := connectDB()
	defer db.Close()
	dsflksd
	handleRequests()

}
