package config

import (
	"encoding/json"

	"github.com/google/wire"
)

var (
	defaultConfig = &Config{
		Log: &Log{
			Level:  "info",
			Output: "console",
		},
		HTTP: &HTTP{
			Host: "0.0.0.0",
			Port: 1992,
			Mode: "debug",
		},
		DB: &DB{
			URL:    "root:changeme@tcp(localhost:3306)/ryze",
			Source: "github://blackhorseya/ryze/scripts/migrations",
		},
		ETH: &ETH{
			URL:       "<url>",
			Websocket: "<websocket>",
		},
		Kafka: &KafkaOptions{
			Brokers:  []string{"localhost:9092"},
			Username: "root",
			Password: "changeme",
		},
	}
)

// ProviderSet is a provider set for wire
var ProviderSet = wire.NewSet(NewConfig)

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
	Conns  int    `json:"conns" yaml:"conns"`
	Source string `json:"source" yaml:"source"`
}

type ETH struct {
	URL       string `json:"url" yaml:"url"`
	Websocket string `json:"websocket" yaml:"websocket"`
}

type KafkaOptions struct {
	Brokers  []string `json:"brokers" yaml:"brokers"`
	Username string   `json:"username" yaml:"username"`
	Password string   `json:"password" yaml:"password"`
}

type Config struct {
	Log   *Log          `json:"log" yaml:"log"`
	HTTP  *HTTP         `json:"http" yaml:"http"`
	DB    *DB           `json:"db" yaml:"db"`
	ETH   *ETH          `json:"eth" yaml:"eth"`
	Kafka *KafkaOptions `json:"kafka" yaml:"kafka"`
}

func (c *Config) String() string {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return ""
	}

	return string(bytes)
}

func NewConfig(path string) *Config {
	config := defaultConfig

	return config

	// todo: 2023/7/2|sean|impl me
}
