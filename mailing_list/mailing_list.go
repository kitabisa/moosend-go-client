package mailing_list

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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

func (c client) GetAllActiveMailingLists(format commons.Format, withStatistics bool, shortBy ShortBy, sortMethod SortMethod) (returnData []MailingList, err error) {
	url := fmt.Sprintf("%s/lists.%s?apikey=%s&WithStatistics=%t&ShortBy=%s&SortMethod=%s", c.BaseURL, format, c.APIKey, withStatistics, shortBy, sortMethod)
	resp, body, err := commons.MakeRequest(c.HTTPClient, http.MethodGet, url, nil)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	var response GetAllActiveMailingListResponse
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

	returnData = response.Context.MailingLists

	return
}
