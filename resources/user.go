package resources

import (
	"time"
)

var UserFields = []string{"about", "can_see_nsfw", "created", "email", "first_name", "full_name", "hide_pledges", "image_url", "is_email_verified", "last_name", "like_count", "thumb_url", "url", "vanity"}

// The Patreon user, which can be both patron and creator.
type User struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		About           string    `json:"about"`
		CanSeeNsfw      string    `json:"can_see_nsfw"`
		Created         time.Time `json:"created"`
		Email           string    `json:"email"`
		FirstName       string    `json:"first_name"`
		FullName        string    `json:"full_name"`
		HidePledges     bool      `json:"hide_pledges"`
		ImageURL        string    `json:"image_url"`
		IsEmailVerified bool      `json:"is_email_verified"`
		LastName        string    `json:"last_name"`
		LikeCount       string    `json:"like_count"`
		// TODO: social_connections
		ThumbURL string `json:"thumb_url"`
		URL      string `json:"url"`
		Vanity   string `json:"vanity"`
	} `json:"attributes"`
	Relationships struct {
		Campaign    *Relationship      `json:"campaign,omitempty"`
		Memberships *RelationshipArray `json:"memberships,omitempty"`
	} `json:"relationships"`
}

// UserResponse wraps Patreon's fetch user API response
type UserResponse struct {
	Data     User     `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
