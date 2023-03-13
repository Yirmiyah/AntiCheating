package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	clientID     = "fd2f39b97bc675101d01"
	clientSecret = "768165bae44eca0f280bfeaa388d4f72249dedc9"
	scope        = "public_repo"
	redirectURI  = "http://localhost:8080/register"
)

func GetGitHubCodeToken() string {
	authURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s", clientID, scope, url.QueryEscape(redirectURI))
	fmt.Printf("Visit the following URL to grant permission to the app:\n%s\n", authURL)

	// Step 2: Set up an HTTP server to receive the authorization code
	server := &http.Server{Addr: ":8080"}
	codeChan := make(chan string)

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			codeChan <- code
			fmt.Fprintf(w, "Authorization code received. You can close this window now.")
		} else {
			fmt.Fprintf(w, "No authorization code received.")
		}
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("HTTP server error:", err)
		}
	}()

	// Step 3: Wait for the authorization code to be received
	var code2 string
	select {
	case code2 = <-codeChan:
		fmt.Println("Authorization code2 received:", code2)
	case <-time.After(10 * time.Minute):
		fmt.Println("Timed out waiting for authorization code2.")
	}

	token, err := exchangeCodeForToken(code2)
	if err != nil {
		fmt.Println("Failed to exchange code2 for token:", err)
		log.Fatal(err)
	}

	// Step 4: Shutdown the HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("HTTP server shutdown error:", err)
	}

	// Step 3: Exchange the authorization code for an access token

	fmt.Println("Access token:", token)

	// Step 4: Use the access token to make requests to the GitHub API
	// ...

	return token

}

func exchangeCodeForToken(code string) (string, error) {
	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("code", code)
	data.Set("redirect_uri", redirectURI)

	resp, err := http.PostForm("https://github.com/login/oauth/access_token", data)

	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// close the response
	resp.Body.Close()

	values, err := url.ParseQuery(string(body))
	if err != nil {
		log.Fatal(err)
	}

	accessToken := values.Get("access_token")
	if accessToken == "" {
		return "", fmt.Errorf("access token not found in response")
	}

	return accessToken, nil
}
