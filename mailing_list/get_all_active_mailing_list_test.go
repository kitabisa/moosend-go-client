package mailing_list

import (
	"fmt"
	"github.com/kitabisa/moosend-go-client/commons"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func (t *MailingListTestSuite) TestClient_GetAllActiveMailingLists_UnmarshalError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"Code": "0"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetAllActiveMailingLists(commons.JSON, false, CreatedOn, ASC)

	assert.NotNil(t.T(), err)
}

func (t *MailingListTestSuite) TestClient_GetAllActiveMailingLists_HttpStatusNotOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"Code": 100}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetAllActiveMailingLists(commons.JSON, false, CreatedOn, ASC)

	assert.NotNil(t.T(), err)
}

func (t *MailingListTestSuite) TestClient_GetAllActiveMailingLists_CodeNotOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := `{
					"Code": 100,
					"Error": "USER_NOT_FOUND",
					"Context": {}
				}`
		fmt.Fprint(w, body)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetAllActiveMailingLists(commons.JSON, false, CreatedOn, ASC)

	assert.NotNil(t.T(), err)
}

func (t *MailingListTestSuite) TestClient_GetAllActiveMailingLists_CodeOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		body := `{
			"Code": 0,
			"Error": "",
			"Context": {
				"MailingLists": [{
					"ID": "04fad8e2-2b35-4302-a887-58f14a1152ab",
					"Name": "Your List Name",
					"ActiveMemberCount": 1024,
					"BouncedMemberCount": 16,
					"RemovedMemberCount": 32,
					"UnsubscribedMemberCount": 24,
					"Status": 0,
					"CustomFieldsDefinition": [{
							"ID": "dd4fb545-ba00-4afe-bc39-5ed2462fd1d3",
							"Name": "City",
							"Context": null,
							"IsRequired": false,
							"Type": 0
						},
						{
							"ID": "ae8500d7-d1df-4118-9a2c-6654a7d6a6db",
							"Name": "Age",
							"Context": null,
							"IsRequired": true,
							"Type": 1
						}
					],
					"CreatedBy": "127.0.0.1",
					"CreatedOn": "/Date(1368710504000+0300)/",
					"UpdatedBy": "127.0.0.1",
					"UpdatedOn": "/Date(1368710923000+0300)/",
					"ImportOperation": {
						"ID": 0,
						"DataHash": "97de8bcb-ca1c-4a4b-bf02-3f639952e093",
						"Mappings": "Some Mappings",
						"EmailNotify": "Some EmailNotify",
						"CreatedOn": "/Date(1400765707431)/",
						"StartedOn": "/Date(1400765707431)/",
						"CompletedOn": "/Date(1400765707432)/",
						"TotalInserted": 0,
						"TotalUpdated": 0,
						"TotalUnsubscribed": 0,
						"TotalInvalid": 0,
						"TotalDuplicate": 0,
						"TotalMembers": 0,
						"Message": "Some Message",
						"Success": false
					}
				}]
			}
		}`
		fmt.Fprint(w, body)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetAllActiveMailingLists(commons.JSON, false, CreatedOn, ASC)

	assert.Nil(t.T(), err)
}

