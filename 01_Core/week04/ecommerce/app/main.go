package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Sauce is hot sauce
type Sauce struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Price       float64 `json:"price"`
	Spice       int     `json:"spice"`
	Size        float64 `json:"size"`
	ImgPath     string  `json:"imgPath"`
	Featured    bool    `json:"featured"`
	Description string  `json:"description"`
	Ingredients string  `json:"ingredients"`
}

var db *sql.DB
var err error

func connectDB() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	pass := os.Getenv("MYSQL_ROOT_PASSWORD")
	s := "root:" + pass + "@tcp(ecommerce_db_1:3306)/ecommerce"
	db, err = sql.Open("mysql", s)
	if err != nil {
		panic(err)
	}
	return db
}

func getItemsFromDB(q string) []Sauce {
	rows, err := db.Query(q)
	if err != nil {
		panic(err)
	}
	var items []Sauce
	for rows.Next() {
		var item Sauce
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Spice, &item.Size, &item.ImgPath, &item.Brand, &item.Featured, &item.Description, &item.Ingredients)
		if err != nil {
			fmt.Println("error scanning rows")
		}
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		fmt.Println("there was an error with the rows")
	}

	return items
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("Home page hit")
}

func getAllSauces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	sauces := getItemsFromDB("select * from sauces;")
	json.NewEncoder(w).Encode(sauces)
}

func getSingleSauce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	id := mux.Vars(r)["id"]
	sauce := getItemsFromDB("select * from sauces where id=" + id)
	json.NewEncoder(w).Encode(sauce)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/sauces", getAllSauces)
	router.HandleFunc("/sauces/{id}", getSingleSauce)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func main() {

	db := connectDB()
	defer db.Close()

	handleRequests()

}
