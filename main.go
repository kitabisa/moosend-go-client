package moosend_go_client

import (
	mailingList "github.com/kitabisa/moosend-go-client/mailing_list"
	"github.com/kitabisa/moosend-go-client/subscriber"
	"github.com/kitabisa/perkakas/v2/httpclient"
	"time"
)

type Options struct {
	BaseURL string
	APIKey string
}

type Client struct {
	MailingList mailingList.Client
	Subscriber subscriber.Client
}

func NewClient(options Options) Client {
	conf := new(httpclient.HttpClientConf)
	conf.BackoffInterval = 2 * time.Millisecond       // 2ms
	conf.MaximumJitterInterval = 5 * time.Millisecond // 5ms
	conf.Timeout = 15000 * time.Millisecond           // 15s
	conf.RetryCount = 3                               // 3 times

	h := httpclient.NewHttpClient(conf)

	return Client{
		MailingList: mailingList.NewClient(options.BaseURL, options.APIKey, h),
		Subscriber: subscriber.NewClient(options.BaseURL, options.APIKey, h),
	}
}

