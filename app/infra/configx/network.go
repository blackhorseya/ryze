package configx

// Network is the network configuration
type Network struct {
	Name    string `json:"name" yaml:"name"`
	Testnet bool   `json:"testnet" yaml:"testnet"`
}
