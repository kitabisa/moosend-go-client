package subscriber

import (
	"fmt"
	"github.com/kitabisa/moosend-go-client/commons"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func (t *SubscriberTestSuite) TestClient_GetSubsriberByEmail_UnmarshalError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"code": "0"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetSubsriberByEmail(commons.JSON, uuid.NewV4().String(), "someemail@email.com")

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_GetSubsriberByEmail_StatusNotOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		body := `{
					"Code": 100,
					"Error": "USER_NOT_FOUND",
					"Context": {}
				}`
		fmt.Fprint(w, body)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetSubsriberByEmail(commons.JSON, uuid.NewV4().String(), "someemail@email.com")

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_GetSubsriberByEmail_StatusOK() {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			  "Code": 0,
			  "Error": "",
			  "Context": {
				"ID": "f582c388-7916-4890-a9ab-b592615ef095",
				"Name": "Paul",
				"Email": "someemail@email.com",
				"CreatedOn": "/Date(1465379934028)/",
				"UpdatedOn": "/Date(1465379934028)/",
				"UnsubscribedOn": null,
				"UnsubscribedFromID": null,
				"SubscribeType": 1,
				"SubscribeMethod": 2,
				"CustomFields": [
				  {
					"CustomFieldID": "728f6774-37ea-4d81-8607-ce8308136760",
					"Name": "Age",
					"Value": "25"
				  },
				  {
					"CustomFieldID": "b6380e04-c3a6-425b-9931-d25897fa4752",
					"Name": "Country",
					"Value": "USA"
				  }
				],
				"RemovedOn": null
			  }
			}
		`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.GetSubsriberByEmail(commons.JSON, uuid.NewV4().String(), "someemail@email.com")

	assert.Nil(t.T(), err)
}


