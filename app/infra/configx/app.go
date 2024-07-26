package configx

// Application is the application configuration
type Application struct {
	Networks map[string]*Network `json:"networks" yaml:"networks"`
}

// Network is the network configuration
type Network struct {
	Name string `json:"name" yaml:"name"`
}
