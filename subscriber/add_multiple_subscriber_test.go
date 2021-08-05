package subscriber

import (
	"fmt"
	"github.com/kitabisa/moosend-go-client/commons"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func (t *SubscriberTestSuite) TestClient_Add_Multiple_Subscriber_UnmarshalError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"code": "0"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.AddMultipleSubscriber(commons.JSON, uuid.NewV4().String(), AddMultipleSubsRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_Add_Multiple_Subscriber_HttpStatusNotOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"code": "0"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.AddMultipleSubscriber(commons.JSON, uuid.NewV4().String(), AddMultipleSubsRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_Add_Multiple_Subscriber_CodeNotOK() {
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
	_, err := client.AddMultipleSubscriber(commons.JSON, uuid.NewV4().String(), AddMultipleSubsRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_Add_Multiple_Subscriber_CodeOK() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
				"Code": 0,
  				"Error": "2 items with invalid email address were ignored: email@email, email2@email",
  				"Context": [
					{
						"ID": "ca506fc5-0967-4756-8848-74e9766bdbdd",
						"Name": "test1Email",
						"Email": "test1@test.com",
						"CreatedOn": "/Date(1465377493907+0100)/",
						"UpdatedOn": "/Date(1465381164389)/",
						"UnsubscribedOn": null,
						"UnsubscribedFromID": null,
      					"SubscribeType": 1,
      					"SubscribeMethod": 1,
      					"CustomFields": [
        					{
          						"CustomFieldID": "42acf2cf-1096-4c80-960b-051791d9a276",
          						"Name": "Country",
          						"Value": "UK"
        					}
      					],
      					"RemovedOn": null
    				},
   					 {
      					"ID": "b751f349-f6b3-4b14-8d75-c37dfafbe40a",
      					"Name": "test2133Email",
      					"Email": "test2133@test.com",
      					"CreatedOn": "/Date(1465377493907+0100)/",
      					"UpdatedOn": "/Date(1465381164389)/",
      					"UnsubscribedOn": null,
      					"UnsubscribedFromID": null,
      					"SubscribeType": 1,
      					"SubscribeMethod": 1,
      					"CustomFields": [
        					{
          						"CustomFieldID": "60d4e2b0-e5ae-4737-9ac5-ce071ab346fb",
								"Name": "Age",
          						"Value": "25"
        					},
        					{
          						"CustomFieldID": "42acf2cf-1096-4c80-960b-051791d9a276",
          						"Name": "Country",
          						"Value": "USA"
        					}
      					],
						"RemovedOn": null
    				}	
  				]
			}
		`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	_, err := client.AddMultipleSubscriber(commons.JSON, uuid.NewV4().String(), AddMultipleSubsRequest{})

	assert.Nil(t.T(), err)
}