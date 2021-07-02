package commons

import (
	"fmt"
	"github.com/kitabisa/perkakas/v2/httpclient"
	"io"
	"io/ioutil"
	"net/http"
)

func MakeRequest(httpClient *httpclient.HttpClient, method, url string, body io.Reader) (resp *http.Response, bodyResp []byte, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	resp, err = httpClient.Client.Do(req)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bodyResp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("[moosend-client] %s", err.Error())
		return
	}

	return
}

