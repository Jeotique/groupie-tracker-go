package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"project/variables"
	"strings"
)

func TokenAccess() (string, error) {
	clientCreds := fmt.Sprintf("%s:%s", variables.ClientID, variables.ClientSecret)
	clientCredsB64 := base64.StdEncoding.EncodeToString([]byte(clientCreds))

	tokenURL := "https://accounts.spotify.com/api/token"
	tokenData := strings.NewReader("grant_type=client_credentials")

	req, err := http.NewRequest("POST", tokenURL, tokenData)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+clientCredsB64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("spotify API returned non-OK status: %d", resp.StatusCode)
	}

	var tokenResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return "", err
	}

	token, ok := tokenResponse["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("invalid token format")
	}
	fmt.Println(token)
	return token, nil
}

func requestSpotifyAPI(url, token string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	return client.Do(req)
}
