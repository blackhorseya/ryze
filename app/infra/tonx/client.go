package tonx

import (
	"fmt"

	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/stretchr/testify/mock"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

const (
	mainnetConfigURL = "https://ton.org/global.config.json"
	testnetConfigURL = "https://ton-blockchain.github.io/testnet-global.config.json"
)

// Options is a struct that represents the options of the API client.
type Options struct {
	Network string `json:"network" yaml:"network"`
}

// APIClient is a struct that represents the API client.
type APIClient struct {
	*ton.APIClient
	mock.Mock
}

// NewAPIClient is a function that creates a new API client.
func NewAPIClient(options Options) (*APIClient, error) {
	configURL := mainnetConfigURL
	if options.Network == "testnet" {
		configURL = testnetConfigURL
	}

	client := liteclient.NewConnectionPool()

	err := client.AddConnectionsFromConfigUrl(contextx.Background(), configURL)
	if err != nil {
		return nil, fmt.Errorf("failed to add connections from config: %w", err)
	}

	return &APIClient{
		APIClient: ton.NewAPIClient(client),
	}, nil
}
