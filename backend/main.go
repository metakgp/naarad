package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func register(res http.ResponseWriter, req *http.Request) {

}

func main() {
	godotenv.Load()

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	if err != nil {
		panic(err)
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	http.HandleFunc("/register", register)
	fmt.Println("Naarad Backend Server running on port : 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
