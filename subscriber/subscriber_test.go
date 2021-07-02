package subscriber

import (
	"github.com/kitabisa/perkakas/v2/httpclient"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type SubscriberTestSuite struct {
	suite.Suite
	HTTPClient *httpclient.HttpClient
}

func TestSubscriberTestSuite(t *testing.T) {
	suite.Run(t, new(SubscriberTestSuite))
}

func (t *SubscriberTestSuite) SetupTest() {
	conf := new(httpclient.HttpClientConf)
	conf.BackoffInterval = 2 * time.Millisecond       // 2ms
	conf.MaximumJitterInterval = 5 * time.Millisecond // 5ms
	conf.Timeout = 15000 * time.Millisecond           // 15s
	conf.RetryCount = 3                               // 3 times

	t.HTTPClient = httpclient.NewHttpClient(conf)
}

