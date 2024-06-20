package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sethvargo/go-password/password"
)

var (
	db *sql.DB
)

var reqBody struct {
	Uname string `json:"uname"`
}

var resStruct struct {
	PassKey string `json:"pswd"`
	Msg     string `json:"msg"`
}

func register(res http.ResponseWriter, req *http.Request) {
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.Uname == "" {
		http.Error(res, "Username cannot be empty", http.StatusBadRequest)
		return
	}

	pswdGen, err := password.Generate(15, 5, 0, false, false)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	query := fmt.Sprintf("INSERT INTO users VALUES('%s', '%s')", reqBody.Uname, pswdGen)
	_, err = db.Query(query)
	if err != nil {
		http.Error(res, err.Error(), http.StatusConflict)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.PassKey = pswdGen
	resStruct.Msg = "User creation success"

	err = json.NewEncoder(res).Encode(&resStruct)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
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

	http.HandleFunc("POST /register", register)
	fmt.Println("Naarad Backend Server running on port : 3333")
	err = http.ListenAndServe(":3333", nil)

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
