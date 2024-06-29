package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

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
	cookie, _ := req.Cookie("heimdall")
	// It won't throw any error.
	// The service will be protected by heimdall
	// Hence if this endpoint is being triggered then
	// It means that cookie has to be present
	tokenString := cookie.Value

	// Get email from JWT
	reqEmail, _ := http.NewRequest("GET", "https://heimdall-api.metakgp.org/validate-jwt", nil)
	reqEmail.Header.Set("Cookie", fmt.Sprintf("heimdall=%s", tokenString))
	client := &http.Client{}
	resp, err := client.Do(reqEmail)
	if err != nil {
		fmt.Println("heimdall/validate-jwt Error: ", err.Error())
		http.Error(res, "Failed to validate Heimdall session", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&jwtValidateResp); err != nil {
		fmt.Println("Heimdall Response | Email Decoder Error: ", err.Error())
		http.Error(res, "Failed to retrieve user email", http.StatusInternalServerError)
		return
	}

	// Generate user credentials
	userEmail := jwtValidateResp.Email
	username := strings.TrimSuffix(userEmail, "@kgpian.iitkgp.ac.in")
	password := PasswordGenerator(pswdSize)

	// Create user using ntfy api
	signupData := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
	req, _ = http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("NTFY User Registration API Error: ", err.Error())
		http.Error(res, "Failed to request user registration for Naarad", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 409 {
		fmt.Println("User already registered")
		http.Error(res, "User already registered", resp.StatusCode)
		return
	} else if resp.StatusCode != 200 {
		http.Error(res, "Failed to register user", resp.StatusCode)
		return
	}

	// Get the userid from sqlite db
	rowD := db.QueryRow(`SELECT id FROM user WHERE user=?`, username)
	if err = rowD.Scan(&userId); err != nil {
		fmt.Println("Database Error | Get User ID: ", err.Error())
		http.Error(res, "Internal Server Error (DB: Fetch UserID)", http.StatusInternalServerError)
		return
	}

	// Provide read-only access for kgp-* channels to the user
	queryGenAccess := fmt.Sprintf(`INSERT INTO user_access VALUES("%s", "kgp-%%", 1, 0, "")`, userId)
	if _, err = db.Exec(queryGenAccess); err != nil {
		fmt.Println("Granting Access Error: ", err.Error())
		http.Error(res, "Internal Server Error (DB: Access Grant)", http.StatusInternalServerError)
		return
	}

	// Sending user credentials over mail
	emailBody := fmt.Sprintf("Here are the credentials to sign in into Naarad.\n<b>Username</b>: %s\n<b>Password</b>: %s", username, password)
	if sent, err := sendMail(userEmail, "Naarad Login Credentials | Metakgp", emailBody); err != nil || !sent {
		fmt.Println("Sending Credentials Error: ", err.Error())
		http.Error(res, "Failed to send user credentials", http.StatusInternalServerError)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.Msg = "User created successfully!"

	if err = json.NewEncoder(res).Encode(&resStruct); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	initMailer()

	passwordSize, err := strconv.Atoi(os.Getenv("PASSWORD_SIZE"))
	if err != nil {
		pswdSize = 18
	} else {
		pswdSize = passwordSize
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

	if err = db.Ping(); err != nil {
		panic(err)
	}

	http.HandleFunc("GET /register", register)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://naarad.metakgp.org", "http://localhost:3000"},
		AllowCredentials: true,
	})
	fmt.Println("Naarad Backend Server running on port : 5173")
	if err = http.ListenAndServe(":5173", c.Handler(http.DefaultServeMux)); err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	}
}
