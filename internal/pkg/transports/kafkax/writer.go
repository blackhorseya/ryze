package kafkax

import (
	"crypto/tls"

	"github.com/blackhorseya/ryze/internal/pkg/config"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

// NewWriter serve caller to get a new Writer instance
func NewWriter(cfg *config.Config) (*kafka.Writer, error) {
	mechanism := plain.Mechanism{
		Username: cfg.Kafka.Username,
		Password: cfg.Kafka.Password,
	}

	w := &kafka.Writer{
		Addr:     kafka.TCP(cfg.Kafka.Brokers...),
		Balancer: &kafka.Hash{},
		Transport: &kafka.Transport{
			SASL: mechanism,
			TLS: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	return w, nil
}

// WriterSet declare the provider set for writer
var WriterSet = wire.NewSet(NewWriter)
