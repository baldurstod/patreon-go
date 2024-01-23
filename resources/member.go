package resources

import (
	"time"
)

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
		Address                *AddressRelationship      `json:"address"`
		Campaign               *CampaignRelationship     `json:"campaign,omitempty"`
		CurrentlyEntitledTiers *TiersRelationship        `json:"currently_entitled_tiers,omitempty"`
		PledgeHistory          *PledgeEventsRelationship `json:"pledge_history,omitempty"`
		User                   *UserRelationship         `json:"user"`
	} `json:"relationships"`
}
