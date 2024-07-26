package configx

// Configuration is the application configuration
type Configuration struct {
	Networks map[string]*Network `json:"networks" yaml:"networks"`
}
