package patreon

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"os"
	"testing"
	"time"
)

type Config struct {
	ClientID            string `json:"client_id"`
	ClientSecret        string `json:"client_secret"`
	RedirectURL         string `json:"redirect_url"`
	AuthCode            string `json:"auth_code"`
	CreatorAccessToken  string `json:"creator_access_token"`
	CreatorRefreshToken string `json:"creator_refresh_token"`
	PatronAccessToken   string `json:"patron_access_token"`
	PatronRefreshToken  string `json:"patron_refresh_token"`
}

var config = loadConfig()
var scopes = []string{"identity", "campaigns", "campaigns.members"}
var patreonConfig = NewPatreonConfig(config.ClientID, config.ClientSecret, config.RedirectURL, scopes)

func loadConfig() Config {
	var config Config
	f, _ := os.Open("./test_config.json")
	json.NewDecoder(f).Decode(&config)
	return config
}

func TestAuthCode(t *testing.T) {
	patreonClient := NewPatreonClient(patreonConfig)
	err := patreonClient.Exchange(config.AuthCode)
	if err != nil {
		t.Fatal(err)
	}

	token := patreonClient.GetToken()
	fmt.Println("Got token:", token.AccessToken, token.RefreshToken)
}

func TestFetchUser(t *testing.T) {
	patreonClient := NewPatreonClient(patreonConfig)
	patreonClient.SetToken(
		&oauth2.Token{
			AccessToken:  config.PatronAccessToken,
			RefreshToken: config.PatronRefreshToken,
			Expiry:       time.Now().Add(30 * 24 * time.Hour),
		},
	)
	userResponse, err := patreonClient.FetchUser(
		WithIncludes("memberships", "campaign"),
		WithFields("user", "first_name", "thumb_url"),
		WithFields("member", "currently_entitled_amount_cents"),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Fetch user:", userResponse)
}

func TestGetMembership(t *testing.T) {
	patreonClient := NewPatreonClient(patreonConfig)
	patreonClient.SetToken(
		&oauth2.Token{
			AccessToken:  config.PatronAccessToken,
			RefreshToken: config.PatronRefreshToken,
			Expiry:       time.Now().Add(30 * 24 * time.Hour),
		},
	)
	member, err := patreonClient.GetMembership()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("TestGetMembership:", member)
}
