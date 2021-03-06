package subscriber

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/kitabisa/moosend-go-client/commons"
	"github.com/kitabisa/perkakas/v2/httpclient"
)

type client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *httpclient.HttpClient
}

func NewClient(baseUrl, apiKey string, httpClient *httpclient.HttpClient) Client {
	return &client{
		BaseURL:    baseUrl,
		APIKey:     apiKey,
		HTTPClient: httpClient,
	}
}

func (c client) GetSubsriberByEmail(format commons.Format, mailingListID, email string) (returnData Subscriber, err error) {
	apiUrl := fmt.Sprintf("%s/subscribers/%s/view.%s?apikey=%s&Email=%s", c.BaseURL, mailingListID, format, c.APIKey, url.QueryEscape(email))
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodGet, apiUrl, nil)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	returnData = response.Context

	return
}

func (c client) GetSubscriberByID(format commons.Format, mailingListID, id string) (returnData Subscriber, err error) {
	apiUrl := fmt.Sprintf("%s/subscribers/%s/find/%s.%s?apikey=%s", c.BaseURL, mailingListID, id, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodGet, apiUrl, nil)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	returnData = response.Context

	return
}

func (c client) AddSubscriber(format commons.Format, mailingListID string, request SubscribeRequest) (returnData Subscriber, err error) {
	payload, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	apiUrl := fmt.Sprintf("%s/subscribers/%s/subscribe.%s?apikey=%s", c.BaseURL, mailingListID, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodPost, apiUrl, bytes.NewReader(payload))
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	returnData = response.Context
	return
}

func (c client) AddMultipleSubscriber(format commons.Format, mailingListID string, request AddMultipleSubsRequest) (returnData []Subscriber, err error) {
	payload, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
	}

	apiUrl := fmt.Sprintf("%s/subscribers/%s/subscribe_many.%s?apikey=%s", c.BaseURL, mailingListID, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodPost, apiUrl, bytes.NewReader(payload))
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response AddMultipleSubsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	returnData = response.Context

	return
}

func (c client) UpdateSubscriber(format commons.Format, mailingListID, subscriberID string, request SubscribeRequest) (returnData Subscriber, err error) {
	payload, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	apiUrl := fmt.Sprintf("%s/subscribers/%s/update/%s.%s?apikey=%s", c.BaseURL, mailingListID, subscriberID, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodPost, apiUrl, bytes.NewReader(payload))
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	returnData = response.Context

	return
}

func (c client) UnsubscribeFromAccount(format commons.Format, request UnsubscribeRequest) (err error) {
	payload, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	apiUrl := fmt.Sprintf("%s/subscribers/unsubscribe.%s?apikey=%s", c.BaseURL, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodPost, apiUrl, bytes.NewReader(payload))
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	return
}

func (c client) UnsubscribeFromMailingList(format commons.Format, mailingListID string, request UnsubscribeRequest) (err error) {
	payload, err := json.Marshal(request)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	apiUrl := fmt.Sprintf("%s/subscribers/%s/unsubscribe.%s?apikey=%s", c.BaseURL, mailingListID, format, c.APIKey)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodPost, apiUrl, bytes.NewReader(payload))
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, err.Error())
		return
	}

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("[moosend-client] %d:%s", resp.StatusCode, "Unknown error")
		return
	}

	if response.Code != 0 {
		err = errors.New(response.Error)
		return
	}

	return
}