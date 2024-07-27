package configx

import (
	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
)

// Application is the application configuration.
type Application struct {
	Name string `json:"name" yaml:"name"`

	HTTP httpx.Options `json:"http" yaml:"http"`

	OTel struct {
		Target string `json:"target" yaml:"target"`
	} `json:"otel" yaml:"otel"`
}
