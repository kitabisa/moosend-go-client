package subscriber

import (
	"fmt"
	"github.com/kitabisa/moosend-go-client/commons"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
)

func (t *SubscriberTestSuite) TestClient_UnsubscribeFromAccount_StatusNotOK_UnmarshalError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"message": "Publish finished successfully"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	err := client.UnsubscribeFromAccount(commons.JSON, UnsubscribeRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_UnsubscribeFromAccount_StatusNotOK_UnmarshalSuccess() {
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
	err := client.UnsubscribeFromAccount(commons.JSON, UnsubscribeRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_UnsubscribeFromAccount_StatusOK_UnmarshalError() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{"Code": "0"}`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	err := client.UnsubscribeFromAccount(commons.JSON, UnsubscribeRequest{})

	assert.NotNil(t.T(), err)
}

func (t *SubscriberTestSuite) TestClient_UnsubscribeFromAccount_StatusOK_UnmarshalSuccess() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `{
			  "Code": 0,
			  "Error": null,
			  "Context": null
			}
		`)
	}))
	defer ts.Close()

	client := NewClient(ts.URL, uuid.NewV4().String(), t.HTTPClient)
	err := client.UnsubscribeFromAccount(commons.JSON, UnsubscribeRequest{})

	assert.Nil(t.T(), err)
}
