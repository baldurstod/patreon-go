package resources

import (
	"time"
)

var WebhookFields = []string{"last_attempted_at", "num_consecutive_times_failed", "paused", "secret", "uri"}

// Webhooks are fired based on events happening on a particular campaign.
type Webhook struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Last_attempted_at            time.Time `json:"last_attempted_at"`
		Num_consecutive_times_failed int       `json:"num_consecutive_times_failed"`
		Paused                       bool      `json:"paused"`
		Secret                       string    `json:"secret"`
		// TODO: triggers
		Uri string `json:"uri"`
	} `json:"attributes"`
	Relationships struct {
		Campaign *Relationship `json:"campaign,omitempty"`
		Client   *Relationship `json:"client,omitempty"`
	} `json:"relationships"`
}
