package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

var (
	db             *sql.DB
	ntfyServerAddr string
)

var reqBody struct {
	Uname   string `json:"uname"`
	PassKey string `json:"pswd"`
}

var resStruct struct {
	Msg string `json:"msg"`
}

func register(res http.ResponseWriter, req *http.Request) {
	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if reqBody.Uname == "" || reqBody.PassKey == "" {
		http.Error(res, "Username cannot be empty", http.StatusBadRequest)
		return
	}

	// Create user using ntfy api
	signupData := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, reqBody.Uname, reqBody.PassKey)
	req, _ = http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		http.Error(res, "User creation Error", resp.StatusCode)
		return
	}
	// Get the userid from sqlite db
	var (
		userId string
	)

	rowD := db.QueryRow(`SELECT id FROM user WHERE user=?`, reqBody.Uname)

	if err = rowD.Scan(&userId); err != nil {
		http.Error(res, "Database error", http.StatusInternalServerError)
		return
	}

	queryGenAccess := fmt.Sprintf(`INSERT INTO user_access VALUES("%s", "%%", 1, 0, "")`, userId)
	_, err = db.Exec(queryGenAccess)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(res, "Access generation error", http.StatusInternalServerError)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.Msg = "User creation success"

	err = json.NewEncoder(res).Encode(&resStruct)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	ntfyServerAddr = os.Getenv("NTFY_SERVER")
	db, err = sql.Open("sqlite3", "user.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("POST /register", register)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"https://naarad.metakgp.org", "http://localhost:3000"},
	})
	fmt.Println("Naarad Backend Server running on port : 3333")
	err = http.ListenAndServe(":3333", c.Handler(http.DefaultServeMux))

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
