package resources

import (
	"time"
)

// A funding goal in USD set by a creator on a campaign.
type Goal struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AmountCents         int       `json:"amount_cents"`
		CompletedPercentage int       `json:"completed_percentage"`
		CreatedAt           time.Time `json:"created_at"`
		Description         string    `json:"description"`
		ReachedAt           time.Time `json:"reached_at"`
		Title               string    `json:"title"`
	} `json:"attributes"`
	Relationships struct {
		Campaign   *Relationship      `json:"campaign,omitempty"`
	} `json:"relationships"`
}
