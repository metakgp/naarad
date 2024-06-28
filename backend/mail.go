package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

const (
	TOKEN_FILE       = "token.json"
	CREDENTIALS_FILE = "credentials.json"
)

// Returns the generated client.
func getClient(config *oauth2.Config) (*http.Client, error) {
	tok, err := tokenFromFile(TOKEN_FILE)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from file: %v", err)
	}
	return config.Client(context.Background(), tok), nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n\n", authURL)

	fmt.Print("Enter the authorization code: ")
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, fmt.Errorf("unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve token from web: %v", err)
	}
	return tok, nil
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
	return nil
}

// Retrieve a token, saves the token
// The file token.json stores the user's access and refresh tokens
func initMailer() error {
	if _, err := os.Stat(TOKEN_FILE); err == nil {
		fmt.Println("Token file already exists. Proceeding")
		return nil
	}
	fmt.Println("Token file not found. Generating new token")

	b, err := os.ReadFile(CREDENTIALS_FILE)
	if err != nil {
		return fmt.Errorf("unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		return fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	token, err := getTokenFromWeb(config)
	if err != nil {
		return fmt.Errorf("unable to retrieve token from web: %v", err)
	}

	err = saveToken(TOKEN_FILE, token)
	if err != nil {
		return fmt.Errorf("unable to save token: %v", err)
	}

	return nil
}

func sendMail(receiverEmail string, subject string, body string) (bool, error) {
	ctx := context.Background()
	b, err := os.ReadFile(CREDENTIALS_FILE)
	if err != nil {
		return false, fmt.Errorf("unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		return false, fmt.Errorf("unable to parse client secret file to config: %v", err)
	}

	client, err := getClient(config)
	if err != nil {
		return false, fmt.Errorf("unable to get client: %v", err)
	}

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return false, fmt.Errorf("unable to retrieve Gmail client: %v", err)
	}

	var message gmail.Message

	msgStr := fmt.Sprintf("From: 'me'\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", receiverEmail, subject, body)
	message.Raw = base64.URLEncoding.EncodeToString([]byte(msgStr))

	userID := "me"
	_, err = srv.Users.Messages.Send(userID, &message).Do()
	if err != nil {
		return false, fmt.Errorf("unable to send message: %v", err)
	}

	return true, nil
}
