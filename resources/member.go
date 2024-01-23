package resources

import (
	"time"
)

var MemberFields = []string{"campaign_lifetime_support_cents", "currently_entitled_amount_cents", "email", "full_name", "is_follower", "last_charge_date", "last_charge_status", "lifetime_support_cents", "next_charge_date", "note", "patron_status", "pledge_cadence", "pledge_relationship_start", "will_pay_amount_cents"}

// The record of a user's membership to a campaign. Remains consistent across months of pledging.
type Member struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		CampaignLifetimeSupportCents       int       `json:"campaign_lifetime_support_cents"`
		CurrentlyEntitledAmountCents       int       `json:"currently_entitled_amount_cents"`
		Email                              string    `json:"email"`
		FullName                           string    `json:"full_name"`
		IsFollower                         bool      `json:"is_follower"`
		LastChargeDate                     time.Time `json:"last_charge_date"`
		LastChargeStatus                   string    `json:"last_charge_status"`
		LifetimeSupportCents               int       `json:"lifetime_support_cents"`
		NextChargeDate                     time.Time `json:"next_charge_date"`
		Note                               string    `json:"note"`
		PatronStatus                       string    `json:"patron_status"`
		PledgeCadence                      int       `json:"pledge_cadence"`
		PledgeRelationshipStart            time.Time `json:"pledge_relationship_start"`
		WillPayAmountCents                 int       `json:"will_pay_amount_cents"`
	} `json:"attributes"`
	Relationships struct {
		Address                *Relationship      `json:"address,omitempty"`
		Campaign               *Relationship      `json:"campaign,omitempty"`
		CurrentlyEntitledTiers *RelationshipArray `json:"currently_entitled_tiers,omitempty"`
		PledgeHistory          *RelationshipArray `json:"pledge_history,omitempty"`
		User                   *Relationship      `json:"user,omitempty"`
	} `json:"relationships"`
}

// MemberResponse wraps Patreon's fetch member API response
type MemberResponse struct {
	Data     Member   `json:"data"`
	Included Includes `json:"included"`
	Links    struct {
		Self string `json:"self"`
	} `json:"links"`
}
