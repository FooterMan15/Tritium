package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
}

func msal() (string, error) {
	clientId := os.Getenv("CLIENT_ID")
	clientScreret := os.Getenv("SECRET")

	data := url.Values{}
	data.Set("client_id", clientId)
	data.Set("client_secret", clientScreret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "XboxLive.signin")

	req, err := http.NewRequest("POST", "https://login.live.com/oauth20_token.srf", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read body: %v", err)
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	fmt.Printf("Token: %s\n", tokenResponse.AccessToken)

	return "access_token", nil
}
