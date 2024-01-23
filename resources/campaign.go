package resources

import (
	"time"
)

var CampaignFields = []string{"created_at", "creation_name", "discord_server_id", "google_analytics_id", "has_rss", "has_sent_rss_notify", "image_small_url", "image_url", "is_charged_immediately", "is_monthly", "is_nsfw", "main_video_embed", "main_video_url", "one_liner", "patron_count", "pay_per_name", "pledge_url", "published_at", "rss_artwork_url", "rss_feed_title", "show_earnings", "summary", "thanks_embed", "thanks_msg", "thanks_video_url", "url", "vanity"}

// The creator's page, and the top-level object for accessing lists of members, tiers, etc.
type Campaign struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CreatedAt                     time.Time `json:"created_at"`
		CreationName                  string    `json:"creation_name"`
		DiscordServerId               string    `json:"discord_server_id"`
		GoogleAnalyticsId             string    `json:"google_analytics_id"`
		HasRss                        bool      `json:"has_rss"`
		HasSentRssNotify              bool      `json:"has_sent_rss_notify"`
		ImageSmallURL                 string    `json:"image_small_url"`
		ImageURL                      string    `json:"image_url"`
		IsChargedImmediately          bool      `json:"is_charged_immediately"`
		IsMonthly                     bool      `json:"is_monthly"`
		IsNsfw                        bool      `json:"is_nsfw"`
		MainVideoEmbed                string    `json:"main_video_embed"`
		MainVideoURL                  string    `json:"main_video_url"`
		OneLiner                      string    `json:"one_liner"`
		PatronCount                   int       `json:"patron_count"`
		PayPerName                    string    `json:"pay_per_name"`
		PledgeURL                     string    `json:"pledge_url"`
		PublishedAt                   time.Time `json:"published_at"`
		RssArtworkUrl                 string    `json:"rss_artwork_url"`
		RssFeedTitle                  string    `json:"rss_feed_title"`
		ShowEarnings                  bool      `json:"show_earnings"`
		Summary                       string    `json:"summary"`
		ThanksEmbed                   string    `json:"thanks_embed"`
		ThanksMsg                     string    `json:"thanks_msg"`
		ThanksVideoURL                string    `json:"thanks_video_url"`
		Url                           string    `json:"url"`
		Vanity                        string    `json:"vanity"`
	} `json:"attributes"`
	Relationships struct {
		Benefits               *RelationshipArray      `json:"benefits,omitempty"`
		CampaignInstallations  *RelationshipArray      `json:"campaign_installations,omitempty"`
		Categories             *RelationshipArray      `json:"categories,omitempty"`
		Creator                *Relationship           `json:"creator,omitempty"`
		Goals                  *RelationshipArray      `json:"goals,omitempty"`
		Tiers                  *RelationshipArray      `json:"tiers,omitempty"`
	} `json:"relationships"`
}

// CampaignResponse wraps Patreon's campaign API response
type CampaignResponse struct {
	Data     []Campaign `json:"data"`
	Included Includes   `json:"included"`
}
