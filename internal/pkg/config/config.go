package config

import (
	"encoding/json"

	"github.com/google/wire"
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewViper)

type Log struct {
	Level  string `json:"level" yaml:"level"`
	Output string `json:"output" yaml:"output"`
}

type HTTP struct {
	Host string `json:"host" yaml:"host"`
	Port int    `json:"port" yaml:"port"`
	Mode string `json:"mode" yaml:"mode"`
}

type DB struct {
	URL    string `json:"url" yaml:"url"`
	Source string `json:"source" yaml:"source"`
}

type Settings struct {
	Log  *Log  `json:"log" yaml:"log"`
	HTTP *HTTP `json:"http" yaml:"http"`
	DB   *DB   `json:"db" yaml:"db"`
}

func (s *Settings) String() string {
	bytes, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return ""
	}

	return string(bytes)
}
