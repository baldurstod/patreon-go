package resources

import (
	"time"
)

var BenefitFields = []string{"app_external_id", "benefit_type", "created_at", "deliverables_due_today_count", "delivered_deliverables_count", "description", "is_deleted", "is_ended", "is_published", "next_deliverable_due_date", "not_delivered_deliverables_count", "rule_type", "tiers_count", "title",}

// A benefit added to the campaign, which can be added to a tier to be delivered to the patron.
type Benefit struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		AppExternalId                 string    `json:"app_external_id"`
		// TODO: app_meta
		BenefitType                   string    `json:"benefit_type"`
		CreatedAt                     time.Time `json:"created_at"`
		DeliverablesDueTodayCount     int       `json:"deliverables_due_today_count"`
		DeliveredDeliverablesCount    int       `json:"delivered_deliverables_count"`
		description                   string    `json:"description"`
		IsDeleted                     bool      `json:"is_deleted"`
		IsEnded                       bool      `json:"is_ended"`
		IsPublished                   bool      `json:"is_published"`
		NextDeliverableDueDate        time.Time `json:"next_deliverable_due_date"`
		NotDeliveredDeliverablesCount int       `json:"not_delivered_deliverables_count"`
		RuleType                      string    `json:"rule_type"`
		TiersCount                    int       `json:"tiers_count"`
		Title                         string    `json:"title"`
	} `json:"attributes"`
	Relationships struct {
		Campaign              *Relationship      `json:"campaign,omitempty"`
		CampaignInstallation  *Relationship      `json:"campaign_installation,omitempty"`
		Deliverables          *RelationshipArray `json:"deliverables,omitempty"`
		Tiers                 *RelationshipArray `json:"tiers,omitempty"`
	} `json:"relationships"`
}
