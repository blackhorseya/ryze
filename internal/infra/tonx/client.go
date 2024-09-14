package tonx

import (
	"fmt"

	"github.com/blackhorseya/ryze/pkg/contextx"
	"github.com/stretchr/testify/mock"
	"github.com/xssnick/tonutils-go/liteclient"
)

const (
	mainnetConfigURL = "https://ton.org/global.config.json"
	testnetConfigURL = "https://ton-blockchain.github.io/testnet-global.config.json"
)

// Options is a struct that represents the options of the API client.
type Options struct {
	Network string `json:"network" yaml:"network"`
}

// Client is a struct that represents the API client.
type Client struct {
	*liteclient.ConnectionPool
	mock.Mock

	Config *liteclient.GlobalConfig
}

// NewClient is a function that creates a new API client.
func NewClient(options Options) (*Client, error) {
	configURL := mainnetConfigURL
	if options.Network == "testnet" {
		configURL = testnetConfigURL
	}

	client := liteclient.NewConnectionPool()

	config, err := liteclient.GetConfigFromUrl(contextx.Background(), configURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get config from url: %w", err)
	}

	err = client.AddConnectionsFromConfig(contextx.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to add connections from config: %w", err)
	}

	return &Client{
		ConnectionPool: client,
		Config:         config,
	}, nil
}
