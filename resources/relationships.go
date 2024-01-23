package resources

// Data represents a link to entity.
type Data struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type RelationshipArray struct {
	Data []Data `json:"data"`
}
type Relationship struct {
	Data  Data `json:"data"`
	Links struct {
		Related string `json:"related"`
	} `json:"links"`
}
