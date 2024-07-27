package configx

import (
	"github.com/blackhorseya/ryze/app/infra/transports/httpx"
)

// Application is the application configuration.
type Application struct {
	HTTP httpx.Options `json:"http" yaml:"http"`
}
