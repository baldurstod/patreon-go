package resources

import (
	"time"
)

var PledgeEventFields = []string{"amount_cents", "currency_code", "date", "payment_status", "tier_id", "tier_title", "type"}

// The record of a pledging action taken by the user, or that action's failure.
type PledgeEvent struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AmountCents   int       `json:"amount_cents"`
		CurrencyCode  string    `json:"currency_code"`
		Date          time.Time `json:"date"`
		PaymentStatus string    `json:"payment_status"`
		TierId        string    `json:"tier_id"`
		TierTitle     string    `json:"tier_title"`
		Type          string    `json:"type"`
	} `json:"attributes"`
	Relationships struct {
		Campaign *Relationship `json:"campaign,omitempty"`
		Patron   *Relationship `json:"patron,omitempty"`
		Tier     *Relationship `json:"tier,omitempty"`
	} `json:"relationships"`
}
