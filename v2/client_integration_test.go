package binance

import (
	"os"
	"testing"

	"github.com/adshao/go-binance/v2/common"
	"github.com/stretchr/testify/suite"
)

type baseIntegrationTestSuite struct {
	suite.Suite
	client *Client
}

func SetupTest(t *testing.T) *baseIntegrationTestSuite {
	apiKey := os.Getenv("BINANCE_API_KEY")
	secretKey := os.Getenv("BINANCE_SECRET_KEY")
	proxyURL := os.Getenv("BINANCE_PROXY_URL")
	useTestnet := true
	if os.Getenv("BINANCE_USE_TESTNET") == "false" {
		useTestnet = false
	}

	if apiKey == "" || secretKey == "" {
		t.Skip("API key and secret are required for integration tests")
	}

	var client *Client
	if proxyURL != "" {
		client = NewProxiedClient(apiKey, secretKey, proxyURL, common.UseTestnet(useTestnet))
	} else {
		client = NewClient(apiKey, secretKey, common.UseTestnet(useTestnet))
	}

	client.Debug = true

	return &baseIntegrationTestSuite{
		client: client,
	}
}
