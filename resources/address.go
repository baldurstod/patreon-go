package resources

import (
	"time"
)

var AddressFields = []string{"addressee", "city", "country", "created_at", "line_1", "line_2", "phone_number", "postal_code", "state"}

// A patron's shipping address.
type Address struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Addressee   string    `json:"addressee"`
		City        string    `json:"city"`
		Country     string    `json:"country"`
		CreatedAt   time.Time `json:"created_at"`
		Line1       string    `json:"line_1"`
		Line2       string    `json:"line_2"`
		PhoneNumber string    `json:"phone_number"`
		PostalCode  string    `json:"postal_code"`
		State       string    `json:"state"`
	} `json:"attributes"`
	Relationships struct {
		Campaigns              *RelationshipArray    `json:"campaigns,omitempty"`
		User                   *Relationship         `json:"user,omitempty"`
	} `json:"relationships"`
}
