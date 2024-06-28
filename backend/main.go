package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

var (
	db             *sql.DB
	ntfyServerAddr string
	userId         string
	pswdSize       int
)

const (
	lowerCase   = "abcdefghijklmnopqrstuvwxyz"
	upperCase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers     = "0123456789"
	specialChar = "!@#$%^&*()_-+{}[]"
)

var resStruct struct {
	Msg string `json:"msg"`
}

var jwtValidateResp struct {
	Email string `json:"email"`
}

func PasswordGenerator(passwordLength int) string {
	password := ""
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	for n := 0; n < passwordLength; n++ {
		randNum := randGen.Intn(4)

		switch randNum {
		case 0:
			randCharNum := randGen.Intn(len(lowerCase))
			password += string(lowerCase[randCharNum])
		case 1:
			randCharNum := randGen.Intn(len(upperCase))
			password += string(upperCase[randCharNum])
		case 2:
			randCharNum := randGen.Intn(len(numbers))
			password += string(numbers[randCharNum])
		case 3:
			randCharNum := randGen.Intn(len(specialChar))
			password += string(specialChar[randCharNum])
		}
	}

	return password
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
		fmt.Println("Heimdall API Error: ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&jwtValidateResp); err != nil {
		fmt.Println("Email Decoder Error: ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	uname := strings.TrimSuffix(jwtValidateResp.Email, "@kgpian.iitkgp.ac.in")
	userEmail := jwtValidateResp.Email
	pswd := PasswordGenerator(pswdSize)

	// Create user using ntfy api
	signupData := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, uname, pswd)
	req, _ = http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))

	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("NTFY User API error: ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
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
		fmt.Println("Database Error: ", err.Error())
		http.Error(res, "Internal server error", http.StatusInternalServerError)
		return
	}

	queryGenAccess := fmt.Sprintf(`INSERT INTO user_access VALUES("%s", "kgp-%%", 1, 0, "")`, userId)
	_, err = db.Exec(queryGenAccess)

	if err != nil {
		fmt.Println("Access generation error: ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	emailBody := fmt.Sprintf("Username for signing in to Naarad portal: %s\nPassword for signing in to Naarad Portal: %s", uname, pswd)
	sent, err := sendMail(userEmail, "Naarad Login Details | Metakgp", emailBody)

	if err != nil || !sent {
		fmt.Println("Error sending confidentials: ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.Msg = "User created successfully"

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

	initMailer()

	ntfyServerAddr = os.Getenv("NTFY_SERVER")
	fileLoc := os.Getenv("NTFY_AUTH_FILE")
	pswdSize, err = strconv.Atoi(os.Getenv("PSWD_SIZE"))
	if err != nil {
		pswdSize = 18
	}

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

	http.HandleFunc("GET /register", register)

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
