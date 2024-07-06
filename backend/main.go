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

var healthResponse struct {
	Healthy bool `json:"healthy"`
}

type responseRecorder struct {
	http.ResponseWriter
	status int
	size   int
}

func (r *responseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &responseRecorder{w, http.StatusOK, 0}
		next.ServeHTTP(recorder, r)

		logLevel := "INFO"
		if recorder.status >= 400 {
			logLevel = "ERROR"
		}

		log.Printf("%s:\t%s - %q %s %d %s\n",
			logLevel,
			r.Header.Get("X-Real-IP"),
			r.Method,
			r.RequestURI,
			recorder.status,
			http.StatusText(recorder.status),
		)
	})
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
		fmt.Println("[ERROR] ~ Retrieving Heimdall Token: ", err.Error())
		http.Error(res, "[ERROR] ~ Retrieving Heimdall Token", http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value

	// Get email from JWT
	reqEmail, _ := http.NewRequest("GET", "https://heimdall-api.metakgp.org/validate-jwt", nil)
	reqEmail.Header.Set("Cookie", fmt.Sprintf("heimdall=%s", tokenString))
	client := &http.Client{}
	resp, err := client.Do(reqEmail)
	if err != nil {
		fmt.Println("[ERROR] ~ Validating Heimdall Token ~", tokenString, ": ", err.Error())
		http.Error(res, "[ERROR] ~ Validating Heimdall Token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&jwtValidateResp); err != nil {
		fmt.Println("[ERROR] ~ Parsing Heimdall Token ~", tokenString, ": ", err.Error())
		http.Error(res, "[ERROR] ~ Parsing Heimdall Token", http.StatusInternalServerError)
		return
	}

	// Generate user credentials
	userEmail := jwtValidateResp.Email
	username := strings.TrimSuffix(userEmail, "@kgpian.iitkgp.ac.in")
	password := PasswordGenerator(pswdSize)

	// Create user using ntfy api
	signupData := fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password)
	reqNtfy, _ := http.NewRequest("POST", ntfyServerAddr+"/v1/account", strings.NewReader(signupData))
	for name, values := range req.Header {
		reqNtfy.Header[name] = values
	}
	resp, err = client.Do(reqNtfy)
	if err != nil {
		fmt.Println("[ERROR] ~ Requesting User Registration ~", username, ": ", err.Error())
		http.Error(res, "[ERROR] ~ Requesting User Registration", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == 409 {
		fmt.Println("[INFO] ~ User Already Registered: ", username)
		http.Error(res, "[INFO] ~ User Already Registered", resp.StatusCode)
		return
	} else if resp.StatusCode != 200 {
		fmt.Println("[ERROR] ~ Registering User: ", username)
		http.Error(res, "[ERROR] ~ Registering User", resp.StatusCode)
		return
	}

	// Get the userid from sqlite db
	rowD := db.QueryRow(`SELECT id FROM user WHERE user=?`, username)
	if err = rowD.Scan(&userId); err != nil {
		fmt.Println("[ERROR] ~ Fetching UserId ~", username, ": ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Provide read-only access for kgp-* channels to the user
	queryGenAccess := fmt.Sprintf(`INSERT INTO user_access VALUES("%s", "kgp-%%", 1, 0, "")`, userId)
	if _, err = db.Exec(queryGenAccess); err != nil {
		fmt.Println("[ERROR] ~ Granting Access to (kgp-*) ~", username, ": ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Provide read-only access for kgp-* channels to the user
	queryGenAccess = fmt.Sprintf(`INSERT INTO user_access VALUES("%s", "st_%%", 1, 0, "")`, userId)
	if _, err = db.Exec(queryGenAccess); err != nil {
		fmt.Println("[ERROR] ~ Granting Access to (st_*) ~", username, ": ", err.Error())
		http.Error(res, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Sending user credentials over mail
	emailBody := fmt.Sprintf("Here are the credentials to sign in into Naarad.\n\nUsername: %s\nPassword: %s", username, password)
	if sent, err := sendMail(userEmail, "Naarad Login Credentials | Metakgp", emailBody); err != nil || !sent {
		fmt.Println("[ERROR] ~ Sending Credentials ~", username, ": ", err.Error())
		http.Error(res, "[ERROR] ~ Sending Credentials", http.StatusInternalServerError)
		return
	}

	http.Header.Add(res.Header(), "content-type", "application/json")
	resStruct.Msg = "[INFO] ~ User " + "(" + username + ")" + "Created Successfully!"

	if err = json.NewEncoder(res).Encode(&resStruct); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func healthCheck(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	// Check database connection
	err := db.Ping()
	if err != nil {
		healthResponse.Healthy = false
	} else {
		healthResponse.Healthy = true
	}

	if !healthResponse.Healthy {
		res.WriteHeader(http.StatusServiceUnavailable)
	} else {
		res.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(res).Encode(healthResponse)
}

func main() {
	godotenv.Load()
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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthCheck)
	mux.HandleFunc("GET /register", register)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://naarad.metakgp.org", "https://naarad-signup.metakgp.org", "http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)
	loggedHandler := LoggerMiddleware(handler)
	if err := http.ListenAndServe(":5173", loggedHandler); err != nil {
		fmt.Printf("error starting server: %s\n", err)
		panic(err)
	} else {
		fmt.Println("Naarad Backend Server running on port : 5173")
	}
}
