package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
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

	signupData := fmt.Sprintf(`{"username": "%s", "pssword": "%s"}`, reqBody.Uname, reqBody.PassKey)
	req, _ = http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))
	req.Header.Set("Accept", "application/json")

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

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.Msg = "User creation success"

	err = json.NewEncoder(res).Encode(&resStruct)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	godotenv.Load()
	ntfyServerAddr = os.Getenv("NTFY_SERVER")
	db, err := sql.Open("sqlite3", "./user.db")
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
