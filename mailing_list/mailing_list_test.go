package mailing_list

import (
	"github.com/kitabisa/perkakas/v2/httpclient"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type MailingListTestSuite struct {
	suite.Suite
	HTTPClient *httpclient.HttpClient
}

func TestMailingListTestSuite(t *testing.T) {
	suite.Run(t, new(MailingListTestSuite))
}

func (t *MailingListTestSuite) SetupTest() {
	conf := new(httpclient.HttpClientConf)
	conf.BackoffInterval = 2 * time.Millisecond       // 2ms
	conf.MaximumJitterInterval = 5 * time.Millisecond // 5ms
	conf.Timeout = 15000 * time.Millisecond           // 15s
	conf.RetryCount = 3                               // 3 times

	t.HTTPClient = httpclient.NewHttpClient(conf)
}

