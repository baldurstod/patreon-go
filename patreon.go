package patreon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/baldurstod/patreon-go/resources"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"strconv"
)

const PatreonAuthURL = "https://www.patreon.com/oauth2/authorize"
const PatreonTokenURL = "https://www.patreon.com/api/oauth2/token"
const PatreonBaseURL = "https://www.patreon.com/"

// PatreonConfig is the configuration of an API client registered here: https://www.patreon.com/portal/registration/register-clients
type PatreonConfig struct {
	clientID     string
	clientSecret string
	redirectURL  string
	scopes       []string
}

// Creates a new PatreonConfig
func NewPatreonConfig(clientID string, clientSecret string, redirectURL string, scopes []string) *PatreonConfig {
	return &PatreonConfig{
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURL:  redirectURL,
		scopes:       scopes,
		//Scopes: []string{"users", "pledges-to-me", "my-campaign"},
	}
}

// PatreonClient is the access point of the patreon API
// You need one PatreonClient per user
type PatreonClient struct {
	httpClient   *http.Client
	oauth2Config *oauth2.Config
	token        *oauth2.Token
}

// Creates a new PatreonClient from a PatreonConfig
// At this point the client is not usable, you need to either:
//   - convert an authorization code into a token with Exchange()
//   - set the token with SetToken()
func NewPatreonClient(patreonConfig *PatreonConfig) *PatreonClient {
	config := oauth2.Config{
		ClientID:     patreonConfig.clientID,
		ClientSecret: patreonConfig.clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  PatreonAuthURL,
			TokenURL: PatreonTokenURL,
		},
		RedirectURL: patreonConfig.redirectURL,
		Scopes:      patreonConfig.scopes,
	}

	return &PatreonClient{
		oauth2Config: &config,
	}
}

// Converts an authorization code into a token.
// The authorization code is available as the code param in the redirected url
func (c *PatreonClient) Exchange(authCode string) error {
	token, err := c.oauth2Config.Exchange(context.Background(), authCode)

	if err == nil {
		c.SetToken(token)
	}

	return err
}

// Set the token associated with this client
// Use case: creating the client from a serialized token (in a session for instance)
func (c *PatreonClient) SetToken(token *oauth2.Token) {
	c.token = token
	c.httpClient = c.oauth2Config.Client(context.Background(), token)
}

// Get the token associated with this client
// Use case: serialisation of the token (in a session for instance)
//   - Note that the token has to be serialized at the end of your process cause it might be refreshed
func (c *PatreonClient) GetToken() *oauth2.Token {
	return c.token
}

func (c *PatreonClient) FetchUser(opts ...requestOption) (*resources.UserResponse, error) {
	resp := &resources.UserResponse{}
	err := c.get("/api/oauth2/v2/identity", resp, opts...)
	return resp, err
}

func (c *PatreonClient) FetchCampaigns(opts ...requestOption) (*resources.CampaignResponse, error) {
	resp := &resources.CampaignResponse{}
	err := c.get("/api/oauth2/v2/campaigns", resp, opts...)
	return resp, err
}

func (c *PatreonClient) FetchMember(id string, opts ...requestOption) (*resources.MemberResponse, error) {
	resp := &resources.MemberResponse{}
	err := c.get("/api/oauth2/v2/members/"+id, resp, opts...)
	return resp, err
}

// Get the membership of the user to your campaign
func (c *PatreonClient) GetMembership() (*resources.Member, error) {
	userResponse, err := c.FetchUser(
		WithIncludes("memberships"),
		WithFields("member", resources.MemberFields...),
	)

	if err != nil {
		return nil, err
	}

	items := userResponse.Included.Items
	for _, item := range items {
		if member, ok := item.(*resources.Member); ok {
			return member, nil
		}
	}

	return nil, errors.New("no membership found")
}

func (c *PatreonClient) buildURL(path string, opts ...requestOption) (string, error) {
	cfg := getOptions(opts...)

	fullURL, _ := url.JoinPath(PatreonBaseURL, path)
	u, err := url.ParseRequestURI(fullURL)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	if cfg.include != "" {
		q.Set("include", cfg.include)
	}

	if len(cfg.fields) > 0 {
		for resource, fields := range cfg.fields {
			key := fmt.Sprintf("fields[%s]", resource)
			q.Set(key, fields)
		}
	}

	if cfg.size != 0 {
		q.Set("page[count]", strconv.Itoa(cfg.size))
	}

	if cfg.cursor != "" {
		q.Set("page[cursor]", cfg.cursor)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

func (c *PatreonClient) get(path string, v interface{}, opts ...requestOption) error {
	if c.httpClient == nil {
		return errors.New("http client is nil")
	}

	addr, err := c.buildURL(path, opts...)

	if err != nil {
		return err
	}

	resp, err := c.httpClient.Get(addr)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errs := ErrorResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&errs); err != nil {
			return err
		}

		return errs
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
