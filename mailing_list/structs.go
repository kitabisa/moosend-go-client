package mailing_list

type CustomFieldDefinition struct {
	Context    string  `json:"Context,omitempty"`
	ID         string  `json:"ID,omitempty"`
	IsRequired bool    `json:"IsRequired,omitempty"`
	Name       string  `json:"Name,omitempty"`
	Type       float64 `json:"Type,omitempty"`
}

type ImportOperation struct {
	CompletedOn       string  `json:"CompletedOn,omitempty"`
	CreatedOn         string  `json:"CreatedOn,omitempty"`
	DataHash          string  `json:"DataHash,omitempty"`
	EmailNotify       string  `json:"EmailNotify,omitempty"`
	ID                float64 `json:"ID,omitempty"`
	Mappings          string  `json:"Mappings,omitempty"`
	Message           string  `json:"Message,omitempty"`
	StartedOn         string  `json:"StartedOn,omitempty"`
	Success           bool    `json:"Success,omitempty"`
	TotalDuplicate    float64 `json:"TotalDuplicate,omitempty"`
	TotalInserted     float64 `json:"TotalInserted,omitempty"`
	TotalInvalid      float64 `json:"TotalInvalid,omitempty"`
	TotalMembers      float64 `json:"TotalMembers,omitempty"`
	TotalUnsubscribed float64 `json:"TotalUnsubscribed,omitempty"`
	TotalUpdated      float64 `json:"TotalUpdated,omitempty"`
}

type MailingList struct {
	ActiveMemberCount       float64                 `json:"ActiveMemberCount,omitempty"`
	BouncedMemberCount      float64                 `json:"BouncedMemberCount,omitempty"`
	CreatedBy               string                  `json:"CreatedBy,omitempty"`
	CreatedOn               string                  `json:"CreatedOn,omitempty"`
	CustomFieldDefinitions  []CustomFieldDefinition `json:"CustomFieldsDefinition,omitempty"`
	ID                      string                  `json:"ID,omitempty"`
	ImportOperation         ImportOperation         `json:"ImportOperation,omitempty"`
	Name                    string                  `json:"Name,omitempty"`
	RemovedMemberCount      float64                 `json:"RemovedMemberCount,omitempty"`
	Status                  float64                 `json:"Status,omitempty"`
	UnsubscribedMemberCount float64                 `json:"UnsubscribedMemberCount,omitempty"`
	UpdatedBy               string                  `json:"UpdatedBy,omitempty"`
	UpdatedOn               string                  `json:"UpdatedOn,omitempty"`
}

type MailingListsContext struct {
	MailingLists []MailingList `json:"MailingLists"`
}

type GetAllActiveMailingListResponse struct {
	Code    float64             `json:"Code"`
	Context MailingListsContext `json:"Context"`
	Error   string              `json:"Error"`
}
