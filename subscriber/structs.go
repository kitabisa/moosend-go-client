package subscriber

type CustomField struct {
	CustomFieldID string `json:"CustomFieldID,omitempty"`
	Name          string `json:"Name,omitempty"`
	Value         string `json:"Value,omitempty"`
}

type Subscriber struct {
	CreatedOn          string        `json:"CreatedOn,omitempty"`
	CustomFields       []CustomField `json:"CustomFields,omitempty"`
	Email              string        `json:"Email,omitempty"`
	ID                 string        `json:"ID,omitempty"`
	Name               string        `json:"Name,omitempty"`
	RemovedOn          string        `json:"RemovedOn,omitempty"`
	SubscribeMethod    float64       `json:"SubscribeMethod,omitempty"`
	SubscribeType      float64       `json:"SubscribeType,omitempty"`
	UnsubscribedFromID string        `json:"UnsubscribedFromID,omitempty"`
	UnsubscribedOn     string        `json:"UnsubscribedOn,omitempty"`
	UpdatedOn          string        `json:"UpdatedOn,omitempty"`
}

type SubscribeRequest struct {
	Email                  string   `json:"Email"`
	Name                   string   `json:"Name,omitempty"`
	HasExternalDoubleOptIn bool     `json:"HasExternalDoubleOptIn,omitempty"`
	CustomFields           []string `json:"CustomFields,omitempty"`
}

type UnsubscribeRequest struct {
	Email string `json:"Email"`
}


type Response struct {
	Code     float64    	`json:"Code,omitempty"`
	Context  Subscriber 	`json:"Context,omitempty"`
	Error    string     	`json:"Error,omitempty"`
}

type AddMultipleSubsResponse struct {
	Code     float64    	`json:"Code,omitempty"`
	Context  []Subscriber 	`json:"Context,omitempty"`
	Error    string     	`json:"Error,omitempty"`
}
