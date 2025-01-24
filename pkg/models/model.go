package models

import "time"

type Response struct {
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type DealsPayload struct {
	Deals []Deal `json:"deals"`
}

type ApiResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Deal  `json:"results"`
}

type Deal struct {
	ID               string       `json:"id"`
	RemoteID         string       `json:"remote_id"`
	CreatedAt        time.Time    `json:"created_at"`
	ModifiedAt       time.Time    `json:"modified_at"`
	Name             string       `json:"name"`
	Description      *string      `json:"description"`
	Amount           int          `json:"amount"`
	Owner            string       `json:"owner"`
	Account          *string      `json:"account"`
	Stage            string       `json:"stage"`
	Status           string       `json:"status"`
	LastActivityAt   time.Time    `json:"last_activity_at"`
	CloseDate        time.Time    `json:"close_date"`
	RemoteCreatedAt  time.Time    `json:"remote_created_at"`
	RemoteWasDeleted bool         `json:"remote_was_deleted"`
	FieldMappings    FieldMapping `json:"field_mappings"`
	RemoteData       *interface{} `json:"remote_data"`
}

type FieldMapping struct {
	OrganizationDefinedTargets  map[string]interface{} `json:"organization_defined_targets"`
	LinkedAccountDefinedTargets map[string]interface{} `json:"linked_account_defined_targets"`
}

type HubSpotDeal struct {
	Properties map[string]interface{} `json:"properties"`
}

type Deals struct {
	ID         string `json:"id"`
	CreatedAt  string `json:"createdat"`
	UpdatedAt  string `json:"updatedat"`
	Properties struct {
		DealName   string `json:"dealname"`
		PipeLine   string `json:"pipeline"`
		DealStage  string `json:"dealstage"`
		Amount     string `json:"amount"`
		CloseDate  string `json:"closedate"`
		DealOwner  string `json:"hubspot_owner_id"`
		DealType   string `json:"dealtype"`
		Priority   string `json:"hs_priority"`
		Contact    string `json:"contact"`
		DealStatus string `json:"deal_status"`
	} `json:"properties"`
}
