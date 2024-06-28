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
	userId         string
)

var resStruct struct {
	Msg string `json:"msg"`
}

var jwtValidateResp struct {
	Email string `json:"email"`
}

func getUsername(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("heimdall")
	if err != nil {
		http.Error(res, "No JWT session token found.", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value

	// Get email from jwt token
	reqEmail, _ := http.NewRequest("GET", "https://heimdall-api.metakgp.org/validate-jwt", nil)
	reqEmail.Header.Set("Cookie", fmt.Sprintf("heimdall=%s", tokenString))
	client := &http.Client{}

	resp, err := client.Do(reqEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&jwtValidateResp); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(res).Encode(&jwtValidateResp)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

}

func register(res http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("heimdall")
	if err != nil {
		http.Error(res, "No JWT session token found.", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value

	// Get email from jwt token
	reqEmail, _ := http.NewRequest("GET", "https://heimdall-api.metakgp.org/validate-jwt", nil)
	reqEmail.Header.Set("Cookie", fmt.Sprintf("heimdall=%s", tokenString))
	client := &http.Client{}

	resp, err := client.Do(reqEmail)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&jwtValidateResp); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	uname := strings.TrimSuffix(jwtValidateResp.Email, "@kgpian.iitkgp.ac.in")

	// Create user using ntfy api
	signupData := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, uname, "")
	req, _ = http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))

	resp, err = client.Do(req)
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
	rowD := db.QueryRow(`SELECT id FROM user WHERE user=?`, uname)

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
	fileLoc := os.Getenv("NTFY_AUTH_FILE")
	if fileLoc == "" || ntfyServerAddr == "" {
		panic("NTFY Server or NTFY auth file location cannot be empty")
	}
	db, err = sql.Open("sqlite3", fileLoc)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("POST /register", register)
	http.HandleFunc("GET /uname", getUsername)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://naarad.metakgp.org", "http://localhost:3000"},
		AllowCredentials: true,
	})
	fmt.Println("Naarad Backend Server running on port : 5173")
	err = http.ListenAndServe(":5173", c.Handler(http.DefaultServeMux))

	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
