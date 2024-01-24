package resources

import (
	"time"
)

var TierFields = []string{"amount_cents", "created_at", "description", "edited_at", "image_url", "patron_count", "post_count", "published", "published_at", "remaining", "requires_shipping", "title", "unpublished_at", "url", "user_limit"}

// A membership level on a campaign, which can have benefits attached to it.
type Tier struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AmountCents int       `json:"amount_cents"`
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		//TODO: discord_role_ids
		EditedAt         time.Time `json:"edited_at"`
		ImageUrl         string    `json:"image_url"`
		PatronCount      int       `json:"patron_count"`
		PostCount        int       `json:"post_count"`
		Published        bool      `json:"published"`
		PublishedAt      time.Time `json:"published_at"`
		Remaining        int       `json:"remaining"`
		RequiresShipping bool      `json:"requires_shipping"`
		Title            string    `json:"title"`
		UnpublishedAt    time.Time `json:"unpublished_at"`
		Url              string    `json:"url"`
		UserLimit        int       `json:"user_limit"`
	} `json:"attributes"`
	Relationships struct {
		Benefits  *RelationshipArray `json:"benefits,omitempty"`
		Campaign  *Relationship      `json:"campaign,omitempty"`
		TierImage *Relationship      `json:"tier_image,omitempty"`
	} `json:"relationships"`
}
